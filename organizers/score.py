"""Score a gateway on a benchmark split.

Usage:
    python organizers/score.py --bench bench/dev --gateway "python starter/gateway.py --detector regex"

The gateway command is called as:
    <gateway> sanitize --in <repo> --out <dir> --map <mapping.json>
    <gateway> restore  --in <dir>  --map <mapping.json> --out <dir>
"""
from __future__ import annotations

import argparse
import ast
import difflib
import hashlib
import json
import math
import os
import re
import shlex
import shutil
import subprocess
import sys
import tempfile
import threading
import time
import warnings
from collections import defaultdict
from concurrent.futures import ThreadPoolExecutor
from pathlib import Path

# plain CredData secrets count once, anything mutated or synthetic counts twice
def weight(mutation: str) -> float:
    return 1.0 if mutation == "none" else 2.0


def leaked(form: str, text: str, background: str = "") -> bool:
    """Exact match for short values; for longer ones any window of >=60% of the value (min 8 chars).

    `background` is the original file with this secret cut out. A window that occurs there too is not
    specific to the secret (shared domain suffix, shared DH parameters, ...) and is ignored.
    """
    if len(form) < 8:
        return form in text and form not in background
    w = max(8, math.ceil(0.6 * len(form)))
    return any(form[i:i + w] in text and form[i:i + w] not in background for i in range(len(form) - w + 1))


COLLATERAL = "all"  # set from --collateral
# CredData leaves some weak/default credentials unlabelled (password: "admin"); masking them is not over-masking
# CredData files carry real names and addresses (ChangeLogs, AUTHORS, commit trailers) that CredData does not
# label; masking them is defensible PII handling, so such lines are not counted as over-masking.
EMAIL = re.compile(r"[\w.+-]+@[\w-]+\.[\w.-]+")
CRED_FILE_CAP = 25  # at most this many over-masked lines per CredData file
CRED_LIKE = re.compile(
    r"(?i)(?:pass|pwd|secret|token|api[_-]?key|auth|credential|private[_-]?key|access[_-]?key|salt)\w*[\"']?\s*(?::=|=>|[:=])"
    r"|://[^\s:/@]+:[^\s@/]+@|^[\s\"']*[A-Za-z0-9+/=]{40,}[\"',]*$")
warnings.filterwarnings("ignore", category=SyntaxWarning)  # benchmark files are full of bad escapes


try:
    import tomllib
except ImportError:  # Python < 3.11
    tomllib = None
try:
    import yaml  # optional: PyYAML
except ImportError:
    yaml = None
NODE = shutil.which("node")  # optional: JS syntax via `node --check`
PARSED_EXT = set()  # filled below: suffixes we can check
PARSERS = ["py", "json"] + (["toml"] if tomllib else []) + (["yaml"] if yaml else []) + (["js"] if NODE else [])


PARSED_EXT.update({".py", ".json"} | ({".toml"} if tomllib else set()) | ({".yml", ".yaml"} if yaml else set())
                  | ({".js", ".mjs", ".cjs"} if NODE else set()))
YAML_LOADER = getattr(yaml, "CSafeLoader", None) or getattr(yaml, "SafeLoader", None) if yaml else None


def _node_ok(text: str, ext: str) -> bool:
    with tempfile.NamedTemporaryFile("w", suffix=ext, delete=False, encoding="utf-8") as f:
        f.write(text)
    try:
        return subprocess.run([NODE, "--check", f.name], capture_output=True, timeout=30).returncode == 0
    finally:
        os.unlink(f.name)


_PARSE_CACHE_FILE = Path(os.environ.get("XDG_CACHE_HOME", Path.home() / ".cache")) / "hackatum-score" / "orig_parse.json"
_parse_cache = None
_parse_lock = threading.Lock()


def orig_parses(path: Path, text: str) -> bool | None:
    """parses() for benchmark originals, cached on disk by content hash (originals never change)."""
    global _parse_cache
    key = f"{path.suffix.lower()}:{','.join(PARSERS)}:" + hashlib.sha1(text.encode("utf-8", "replace")).hexdigest()
    with _parse_lock:
        if _parse_cache is None:
            try:
                _parse_cache = json.loads(_PARSE_CACHE_FILE.read_text())
            except (OSError, ValueError):
                _parse_cache = {}
        if key in _parse_cache:
            return _parse_cache[key]
    res = parses(path, text)
    with _parse_lock:
        _parse_cache[key] = res
    return res


def save_parse_cache():
    if _parse_cache:
        try:
            _PARSE_CACHE_FILE.parent.mkdir(parents=True, exist_ok=True)
            tmp = _PARSE_CACHE_FILE.with_suffix(f".{os.getpid()}.tmp")
            tmp.write_text(json.dumps(_parse_cache))
            tmp.replace(_PARSE_CACHE_FILE)
        except OSError:
            pass


def parses(path: Path, text: str) -> bool | None:
    """True/False if we can check this file type, None otherwise. Optional parsers widen the check."""
    ext = path.suffix.lower()
    try:
        if ext == ".py":
            ast.parse(text)
        elif ext == ".json":
            json.loads(text)
        elif ext == ".toml" and tomllib:
            tomllib.loads(text)
        elif ext in (".yml", ".yaml") and yaml:
            list(yaml.load_all(text, Loader=YAML_LOADER))
        elif ext in (".js", ".mjs", ".cjs") and NODE:
            return _node_ok(text, ext)
        else:
            return None
        return True
    except (SyntaxError, ValueError, RecursionError) as e:  # tomllib/json errors subclass ValueError
        return False
    except Exception:  # noqa: BLE001  yaml.YAMLError and friends
        return False


def read(p: Path) -> str:
    return p.read_bytes().decode("utf-8", errors="replace") if p.exists() else ""


def run(cmd: list, timeout: float) -> tuple[bool, float, str]:
    t = time.time()
    try:
        r = subprocess.run(cmd, capture_output=True, text=True, timeout=timeout)
        return r.returncode == 0, time.time() - t, r.stderr[-500:]
    except subprocess.TimeoutExpired:
        return False, time.time() - t, "timeout"


def _occurrences(line: str, values) -> list:
    spans = []
    for v in values:
        start = line.find(v)
        while v and start >= 0:
            spans.append((start, start + len(v)))
            start = line.find(v, start + 1)
    return spans


def collateral(repo: Path, out: Path, labels: dict, scope: str = "synth") -> tuple[int, int]:
    """Masked segments that overlap no secret and no decoy -> (in synthetic files, in CredData files).

    Synthetic files are fully known, so every such segment is over-masking. CredData files are labelled only
    on scanner candidate lines; with scope="all" a masked segment on a line that carries no label at all
    (and is not a multi-line secret) counts too. Lines with a label are judged by that label only.
    """
    synth_files = {s["file"] for s in labels["secrets"] if s.get("source") == "synthetic"}
    files = set(synth_files)
    labelled = defaultdict(set)  # CredData file -> line numbers that carry any label
    if scope == "all":
        for x in labels["secrets"] + labels["decoys"]:
            if x["file"] not in synth_files:
                files.add(x["file"])
                for ln in range(x["line"], x.get("line_end", x["line"]) + 1):
                    labelled[x["file"]].add(ln)
        files |= {str(p.relative_to(repo)) for p in repo.rglob("*") if p.is_file()} - synth_files
    forms = defaultdict(list)
    for s in labels["secrets"]:
        forms[s["file"]] += s["forms"]
    decoy_vals = defaultdict(list)
    for d in labels["decoys"]:
        if "value" in d:
            decoy_vals[d["file"]].append(d["value"])
    count = {True: 0, False: 0}
    for f in files:
        synth = f in synth_files
        before = count[synth]
        orig, san = read(repo / f).split("\n"), read(out / f).split("\n")
        if len(orig) != len(san):
            continue  # already punished: all decoys of this file count as damaged
        for ln, (a, b) in enumerate(zip(orig, san), 1):
            if a == b or (not synth and (ln in labelled[f] or CRED_LIKE.search(a) or EMAIL.search(a))):
                continue  # unlabelled but credential- or person-shaped CredData lines are ambiguous, like X rows
            known = _occurrences(a, forms[f]) + _occurrences(a, decoy_vals[f])
            # autojunk=False is exact but quadratic: a 711k-char line hung the scorer for an hour
            bad = sum(1 for op, i1, i2, _, _ in difflib.SequenceMatcher(None, a, b, autojunk=len(a) > 5000).get_opcodes()
                      if op in ("replace", "delete") and not any(s < i2 and i1 < e for s, e in known))
            # synthetic: every segment; CredData: at most one per line, so one minified file can't dominate
            count[synth] += bad if synth else min(bad, 1)
        if not synth:  # one big data file (a ChangeLog, a CSV) must not decide a repo's precision
            count[False] = min(count[False], before + CRED_FILE_CAP)
    return count[True], count[False]


def score_repo(repo: Path, labels: dict, gateway: list, timeout: float, work: Path) -> dict:
    out, restored, mp = work / "out", work / "restored", work / "map.json"
    ok, dt, err = run(gateway + ["sanitize", "--in", str(repo), "--out", str(out), "--map", str(mp)], timeout)
    res = {"repo": repo.name, "time": round(dt, 2), "ok": ok, "error": "" if ok else err}
    files = [p.relative_to(repo) for p in repo.rglob("*") if p.is_file()]

    # --- leaks
    secrets = []
    cache, orig = {}, {}
    for s in labels["secrets"]:
        if s["file"] not in cache:
            cache[s["file"]] = read(out / s["file"]) if ok else None
            orig[s["file"]] = read(repo / s["file"])
        text = cache[s["file"]]
        if text is None:
            leak = True
        else:
            background = orig[s["file"]]
            for f in sorted(s["forms"], key=len, reverse=True):
                background = background.replace(f, "\0")
            leak = any(leaked(f, text, background) for f in s["forms"])
        secrets.append({"mutation": s["mutation"], "category": s["category"], "leaked": leak,
                        "source": s.get("source", "creddata")})
    res["secrets"] = secrets

    # --- decoys (over-masking)
    decoys_ok = []
    lines_cache = {}
    for d in labels["decoys"]:
        if not ok:
            decoys_ok.append(False)
            continue
        f = d["file"]
        if f not in lines_cache:
            orig_n = read(repo / f).count("\n")
            t = read(out / f)
            lines_cache[f] = t.split("\n") if t.count("\n") == orig_n else None
        lines = lines_cache[f]
        if lines is None or d["line"] > len(lines):
            decoys_ok.append(False)
            continue
        line = lines[d["line"] - 1]
        decoys_ok.append(d["value"] in line if "value" in d else line == d["whole_line"])
    sources = [d.get("source", "creddata") for d in labels["decoys"]]
    # Synthetic files are fully known: every masked segment that touches neither a secret nor a labelled
    # decoy is collateral over-masking and counts as one more damaged decoy.
    if ok:
        n_synth, n_cred = collateral(repo, out, labels, COLLATERAL)
        decoys_ok += [False] * (n_synth + n_cred)
        sources += ["collateral"] * n_synth + ["collateral_creddata"] * n_cred
    res["decoy_sources"] = sources
    res["decoys_intact"] = decoys_ok

    # --- integrity
    if ok:
        mapping = json.loads(read(mp) or "{}")
        by_value = defaultdict(set)
        for ph, v in mapping.items():
            by_value[v].add(ph)
        res["consistency"] = (sum(1 for v in by_value.values() if len(v) == 1) / len(by_value)) if by_value else 1.0
        rok, _, rerr = run(gateway + ["restore", "--in", str(out), "--map", str(mp), "--out", str(restored)], timeout)
        same = [rok and (restored / f).exists() and (restored / f).read_bytes() == (repo / f).read_bytes() for f in files]
        res["roundtrip"] = sum(same) / len(same) if same else 1.0
        # a placeholder standing where a literal was is fine; we check the rest of the structure survived
        ph_re = re.compile("|".join(re.escape(k) for k in sorted(mapping, key=len, reverse=True))) if mapping else None
        parse = []
        for f in files:
            if (repo / f).suffix.lower() not in PARSED_EXT:
                continue
            orig_text, text = read(repo / f), read(out / f)
            if orig_parses(repo / f, orig_text):
                # an untouched file parses exactly like the original: skip the second parser run
                parse.append(True if orig_text == text else bool(parses(out / f, ph_re.sub("0", text) if ph_re else text)))
        res["parse"] = sum(parse) / len(parse) if parse else 1.0
        if not rok:
            res["error"] = "restore failed: " + rerr
    else:
        res.update(consistency=0.0, roundtrip=0.0, parse=0.0)
    return res


def aggregate(results: list, beta: float = 1.0) -> dict:
    caught_w = total_w = 0.0
    by_mut = defaultdict(lambda: [0, 0])
    by_cat = defaultdict(lambda: [0, 0])
    for r in results:
        for s in r["secrets"]:
            w = weight(s["mutation"])
            total_w += w
            caught_w += 0 if s["leaked"] else w
            by_mut[s["mutation"]][0] += 0 if s["leaked"] else 1
            by_mut[s["mutation"]][1] += 1
            cat = s["category"].split(":")[0]
            by_cat[cat][0] += 0 if s["leaked"] else 1
            by_cat[cat][1] += 1
    decoys = [d for r in results for d in r["decoys_intact"]]
    by_src_decoy = defaultdict(lambda: [0, 0])
    by_src_secret = defaultdict(lambda: [0, 0])
    for r in results:
        for ok, src in zip(r["decoys_intact"], r["decoy_sources"]):
            by_src_decoy[src][0] += ok
            by_src_decoy[src][1] += 1
        for s in r["secrets"]:
            by_src_secret[s["source"]][0] += not s["leaked"]
            by_src_secret[s["source"]][1] += 1
    recall = caught_w / total_w if total_w else 1.0
    precision = sum(decoys) / len(decoys) if decoys else 1.0
    def fbeta(b):
        return (1 + b * b) * precision * recall / (b * b * precision + recall) if precision + recall else 0.0
    f = fbeta(beta)
    n = len(results)
    integrity = (sum(r["roundtrip"] for r in results) / n) * (sum(r["parse"] for r in results) / n) \
        * (0.5 + 0.5 * sum(r["consistency"] for r in results) / n)
    return {
        "score": round(100 * f * integrity, 2),
        "score_f1": round(100 * fbeta(1) * integrity, 2),
        "score_f2": round(100 * fbeta(2) * integrity, 2),
        "recall_weighted": round(recall, 4),
        "decoy_precision": round(precision, 4),
        "beta": beta,
        "integrity": round(integrity, 4),
        "recall_by_source": {k: f"{c}/{t} = {c / t:.2%}" for k, (c, t) in sorted(by_src_secret.items())},
        "decoy_precision_by_source": {k: f"{c}/{t} = {c / t:.2%}" for k, (c, t) in sorted(by_src_decoy.items())},
        "recall_by_mutation": {k: f"{c}/{t} = {c / t:.2%}" for k, (c, t) in sorted(by_mut.items())},
        "recall_by_category": {k: f"{c}/{t} = {c / t:.2%}" for k, (c, t) in
                               sorted(by_cat.items(), key=lambda x: -x[1][1])[:16]},
        "repos": n,
        "failed_repos": [r["repo"] for r in results if not r["ok"]],
        "total_time_s": round(sum(r["time"] for r in results), 1),
    }


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--bench", required=True, type=Path)
    ap.add_argument("--gateway", required=True, help="command to run the gateway")
    ap.add_argument("--timeout", default=300.0, type=float, help="seconds per repo per command (the announced limit: 300)")
    ap.add_argument("--report", type=Path, help="write per-repo details as json")
    ap.add_argument("--jobs", default=1, type=int, help="repos scored in parallel")
    ap.add_argument("--beta", default=1.0, type=float, help="F-beta of recall vs decoy precision (score_f1/score_f2 always shown)")
    ap.add_argument("--collateral", default="all", choices=["synth", "all"],
                    help="where masking unlabelled text counts as a damaged decoy (all = also CredData files)")
    ap.add_argument("--sample", type=int, default=0,
                    help="score only N repos picked at random (same N and --seed give the same repos); 0 = all")
    ap.add_argument("--seed", type=int, default=0, help="seed for --sample")
    args = ap.parse_args()
    global COLLATERAL
    COLLATERAL = args.collateral
    gateway = shlex.split(args.gateway)
    def one(lab_path):
        repo = args.bench / lab_path.name.removesuffix(".labels.json")
        with tempfile.TemporaryDirectory() as tmp:
            r = score_repo(repo, json.loads(lab_path.read_text()), gateway, args.timeout, Path(tmp))
        print(f"  {repo.name}: {'ok' if r['ok'] else 'FAIL'} {r['time']}s", file=sys.stderr)
        return r

    labels = sorted(args.bench.glob("*.labels.json"))
    if 0 < args.sample < len(labels):
        import random
        labels = sorted(random.Random(args.seed).sample(labels, args.sample))
        print(f"sample: {len(labels)} repos (--seed {args.seed})", file=sys.stderr)
    with ThreadPoolExecutor(args.jobs) as pool:
        results = list(pool.map(one, labels))
    save_parse_cache()
    summary = aggregate(results, args.beta)
    summary["parse_checked"] = PARSERS
    print(json.dumps(summary, indent=1, ensure_ascii=False))
    if args.report:
        args.report.write_text(json.dumps({"summary": summary, "repos": results}, indent=1))


if __name__ == "__main__":
    main()

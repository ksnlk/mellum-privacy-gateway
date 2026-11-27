"""Detectors: (repo_dir, {rel_path: text}) -> {rel_path: [(start, end), ...]} char offsets to mask.

noop        nothing (lower bound for recall)
regex       known token formats + keyword assignments + URL credentials
entropy     every long high-entropy token (shows the over-masking trap)
mellum      regex/entropy candidates filtered by Mellum2 via LiteLLM (needs MELLUM_KEY)

Off-the-shelf secret scanners are not allowed (see the task README): heuristics propose, Mellum2 decides.
"""
from __future__ import annotations

import json
import math
import os
import re
import sys
import time
import urllib.error
import urllib.request
from collections import Counter
from concurrent.futures import ThreadPoolExecutor
from pathlib import Path

KNOWN = [
    r"AKIA[0-9A-Z]{16}",
    r"gh[pousr]_[A-Za-z0-9]{30,}",
    r"github_pat_[A-Za-z0-9_]{30,}",
    r"xox[abprs]-[A-Za-z0-9-]{10,}",
    r"sk_(?:live|test)_[A-Za-z0-9]{10,}",
    r"AIza[0-9A-Za-z_\-]{30,}",
    r"glpat-[A-Za-z0-9_\-]{20}",
    r"eyJ[A-Za-z0-9_\-]{8,}\.[A-Za-z0-9_\-]{8,}\.[A-Za-z0-9_\-]{8,}",
]
KNOWN_RE = re.compile("|".join(KNOWN))
KEYWORD_RE = re.compile(
    r"(?i)(?:password|passwd|pwd|pass|secret|token|api[_\-]?key|apikey|auth|credential|private[_\-]?key|"
    r"access[_\-]?key|client[_\-]?secret|salt|nonce)[\w\-]*[\"']?\s*(?::=|=>|[:=])\s*[\"']?([^\s\"',;`)\]}]{4,})")
URL_CRED_RE = re.compile(r"[a-z][a-z0-9+.\-]*://[^\s:/@\"']+:([^\s@/\"']{3,})@", re.I)
PEM_RE = re.compile(r"-----BEGIN [A-Z ]*PRIVATE KEY-----\n(.*?)\n-----END", re.S)
TOKEN_RE = re.compile(r"[A-Za-z0-9+/=_\-]{16,}")
PLACEHOLDER_LIKE = re.compile(r"^(?:<.*>|\$\{.*\}|\{\{.*\}\}|%\(.*|x+|\*+|\.\.\.|null|none|true|false|"
                              r"changeme|your[_\-].*|example.*)$", re.I)


def entropy(s: str) -> float:
    c = Counter(s)
    return -sum(n / len(s) * math.log2(n / len(s)) for n in c.values())


CODE_EXT = {".py", ".js", ".ts", ".jsx", ".tsx", ".java", ".go", ".rb", ".php", ".cs", ".kt", ".scala",
            ".swift", ".rs", ".c", ".cc", ".cpp", ".h", ".json"}


def regex_spans(text: str, rel: str = ""):
    in_code = Path(rel).suffix.lower() in CODE_EXT
    spans = [m.span() for m in KNOWN_RE.finditer(text)]
    for rx in (KEYWORD_RE, URL_CRED_RE):
        for m in rx.finditer(text):
            v = m.group(1)
            if "(" in v or v.endswith("\\") or PLACEHOLDER_LIKE.match(v):
                continue
            # in code an unquoted value is an identifier, not a literal
            if rx is KEYWORD_RE and in_code and text[m.start(1) - 1] not in "\"'":
                continue
            spans.append(m.span(1))
    for m in PEM_RE.finditer(text):  # mask each body line separately (no newlines inside a placeholder)
        off = m.start(1)
        for line in m.group(1).split("\n"):
            if line.strip():
                spans.append((off, off + len(line)))
            off += len(line) + 1
    return spans


def entropy_spans(text: str, threshold=3.5):
    return [m.span() for m in TOKEN_RE.finditer(text) if entropy(m.group(0)) > threshold]


def noop(repo, files):
    return {}


def regex(repo, files):
    return {rel: regex_spans(t, rel) for rel, t in files.items()}


def entropy_all(repo, files):
    return {rel: regex_spans(t, rel) + entropy_spans(t) for rel, t in files.items()}


MELLUM_PROMPT = """You are a secret detector. For each numbered candidate answer SECRET or SAFE, one per line,
format `N: SECRET` or `N: SAFE`. A secret is a real hardcoded credential: password, token, API key,
private key material, credentials in a URL. Hashes, ids, versions, placeholders, example values and
variable names are SAFE.

"""


def ask_mellum(content, url, key, model, deadline=None):
    """One chat request. Retries on 429 (key limits) and network errors; None when Mellum never answered
    (or the deadline, a time.monotonic() value, passed)."""
    body = json.dumps({"model": model, "temperature": 0, "max_tokens": 12000,  # it thinks first: ~250 tokens/candidate
                       "messages": [{"role": "user", "content": content}]}).encode()
    req = urllib.request.Request(url, body, {"Content-Type": "application/json", "Authorization": f"Bearer {key}"})
    for attempt in range(int(os.environ.get("MELLUM_RETRIES", 8))):
        left = 600 if deadline is None else deadline - time.monotonic()
        if left < 5:
            return None
        try:
            msg = json.loads(urllib.request.urlopen(req, timeout=min(600, left)).read())["choices"][0]["message"]
            if msg.get("content"):
                return msg["content"]
            wait = 1  # empty answer: thinking ran out of max_tokens, ask again
        except urllib.error.HTTPError as e:
            wait = float(e.headers.get("retry-after") or 5) if e.code == 429 else 2 + 4 * attempt
        except Exception:  # noqa: BLE001  timeouts, resets
            wait = 2 + 4 * attempt
        print(f"mellum retry {attempt}: waiting {wait:.0f}s", file=sys.stderr)
        time.sleep(max(0, min(wait, 60, (600 if deadline is None else deadline - time.monotonic()))))
    return None


def mellum(repo, files, batch=20):
    """Regex+entropy candidates, then Mellum2 decides. One chat request per batch of candidates.

    Env: MELLUM_URL (LiteLLM base url), MELLUM_KEY, MELLUM_MODEL, MELLUM_RETRIES,
    MELLUM_THREADS (requests in flight; default 2 = max_parallel_requests of a team key, see x-ratelimit-* headers),
    GATEWAY_BUDGET_S (seconds for Mellum per repo, default 240; the scorer kills a gateway at its --timeout).
    Fails closed: if Mellum gives no verdict for a batch (error, or the budget ran out), every candidate in it is masked.
    """
    deadline = time.monotonic() + float(os.environ.get("GATEWAY_BUDGET_S", 240))
    url = os.environ.get("MELLUM_URL", "https://tum-hackathon-api.labs.jb.gg").rstrip("/") + "/v1/chat/completions"
    key = os.environ["MELLUM_KEY"]
    model = os.environ.get("MELLUM_MODEL", "JetBrains/Mellum2.1-12B-A2.5B-Thinking")
    jobs = []
    for rel, text in files.items():
        cands = sorted(set(regex_spans(text, rel) + entropy_spans(text)))
        for i in range(0, len(cands), batch):
            jobs.append((rel, text, cands[i:i + batch]))

    def run(job):
        rel, text, chunk = job
        items = []
        for n, (s, e) in enumerate(chunk, 1):
            ls = text.rfind("\n", 0, s) + 1
            le = text.find("\n", e)
            line = text[ls:le if le >= 0 else None].strip()[:200]
            items.append(f"{n}: value {text[s:e][:120]!r} in line: {line}")
        if time.monotonic() > deadline - 5:
            return rel, chunk
        answer = ask_mellum(MELLUM_PROMPT + f"file: {rel}\n" + "\n".join(items), url, key, model, deadline)
        if answer is None:  # no verdict: masking costs a few decoys, skipping would leak
            return rel, chunk
        return rel, [chunk[int(m.group(1)) - 1] for m in re.finditer(r"(\d+)\s*:\s*SECRET", answer, re.I)
                     if 1 <= int(m.group(1)) <= len(chunk)]

    out = {}
    with ThreadPoolExecutor(int(os.environ.get("MELLUM_THREADS", 2))) as pool:
        for rel, keep in pool.map(run, jobs):
            out.setdefault(rel, []).extend(keep)
    return out


REGISTRY = {"noop": noop, "regex": regex, "entropy": entropy_all, "mellum": mellum}

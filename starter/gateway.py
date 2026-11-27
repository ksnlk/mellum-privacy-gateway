"""Starter gateway. Participants replace the detector with something smarter (Mellum2).

    python starter/gateway.py --detector regex sanitize --in repo/ --out sanitized/ --map mapping.json
    python starter/gateway.py restore --in sanitized/ --map mapping.json --out restored/

Contract checked by the scorer:
  * every input file must appear in --out (same relative path)
  * secrets are replaced inline; do not add or remove lines
  * the same secret value gets the same placeholder across the whole repo
  * restore(sanitize(repo)) must give back the original bytes
"""
from __future__ import annotations

import argparse
import json
import re
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent))
import detectors  # noqa: E402

PLACEHOLDER = re.compile(r"<SECRET_(\d+)>")


def read(p: Path) -> str:
    return p.read_bytes().decode("utf-8", errors="surrogateescape")


def write(p: Path, text: str):
    p.parent.mkdir(parents=True, exist_ok=True)
    p.write_bytes(text.encode("utf-8", errors="surrogateescape"))


def merge(spans):
    out = []
    for s, e in sorted(spans):
        if out and s <= out[-1][1]:
            out[-1][1] = max(out[-1][1], e)
        else:
            out.append([s, e])
    return out


def sanitize(src: Path, dst: Path, map_path: Path, detector: str):
    files = {str(p.relative_to(src)): read(p) for p in sorted(src.rglob("*")) if p.is_file()}
    spans = detectors.REGISTRY[detector](src, files)
    value_to_ph, mapping = {}, {}
    for rel, text in files.items():
        pieces, last = [], 0
        for s, e in merge(spans.get(rel, [])):
            value = text[s:e]
            if "\n" in value or not value:
                continue  # keep line structure intact
            if value not in value_to_ph:
                ph = f"<SECRET_{len(value_to_ph) + 1}>"
                value_to_ph[value] = ph
                mapping[ph] = value
            pieces += [text[last:s], value_to_ph[value]]
            last = e
        pieces.append(text[last:])
        write(dst / rel, "".join(pieces))
    map_path.write_text(json.dumps(mapping, indent=1, ensure_ascii=False), encoding="utf-8")


def restore(src: Path, dst: Path, map_path: Path):
    mapping = json.loads(map_path.read_text(encoding="utf-8"))
    for p in sorted(src.rglob("*")):
        if p.is_file():
            text = PLACEHOLDER.sub(lambda m: mapping.get(m.group(0), m.group(0)), read(p))
            write(dst / p.relative_to(src), text)


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--detector", default="regex", choices=sorted(detectors.REGISTRY))
    sub = ap.add_subparsers(dest="cmd", required=True)
    for name in ("sanitize", "restore"):
        p = sub.add_parser(name)
        p.add_argument("--in", dest="src", required=True, type=Path)
        p.add_argument("--out", dest="dst", required=True, type=Path)
        p.add_argument("--map", dest="map", required=True, type=Path)
    a = ap.parse_args()
    if a.cmd == "sanitize":
        sanitize(a.src, a.dst, a.map, a.detector)
    else:
        restore(a.src, a.dst, a.map)


if __name__ == "__main__":
    main()

# Mellum2 Privacy Gateway: notes for coding agents

The task is in `README.md`; it is the source of truth, read it before changing code. This file repeats
the rules an agent breaks most easily, adds practical facts, and lists the commands for the dev loop.
If this file and the README disagree, the README wins.

## Hard rules (breaking them disqualifies the team or zeroes the score)

- **The gateway may call only Mellum2.** No other LLMs, no trained models of any kind: no embeddings,
  no NER (spaCy and the like), nothing with learned weights, also nothing you train yourself.
  Coding agents (you) are fine for *writing* the gateway; the gateway must not call them at runtime.
- **No off-the-shelf secret scanners**, neither as a dependency nor vendored: CredSweeper,
  detect-secrets, gitleaks, trufflehog, semgrep secrets rules, Presidio and the like. Do not install
  them and do not copy their rule files. Hand-written regexes, entropy, parsers, decoders, allow/deny
  lists and thresholds you chose yourself are fine. The organizers read every Dockerfile and dependency list.
- **Never hard-code the key.** Read it from `MELLUM_KEY` and the URL from `MELLUM_URL` if set. Scoring
  runs the image with the organizers' key and `HTTPS_PROXY`, so a baked-in key does not work anyway.
  Never write the key into code, Dockerfile, logs, commits or chat. `.env` is git-ignored; keep it that way.
- **Do not fit to dev values.** The hidden set has fresh secret values and *different* disguises,
  phrasings and file layouts. Never copy values, file names or repo ids from `bench/dev/*.labels.json`
  into the detector. Use the labels to find *classes* of misses, then fix the class.

## Where things live

- `starter/`: the reference gateway. Prefer working in a copy (e.g. `gateway/`) so the baseline stays
  available for comparison; the Dockerfile path is part of the submission, the build context is its directory.
- `organizers/score.py`: the scorer used for the final run. Do not modify it.
- `bench/dev/`: data and labels. Read-only. `bench/licenses/`: upstream licenses, not data.

## Gateway contract (checked by `organizers/score.py`)

`sanitize --in DIR --out DIR --map mapping.json` and `restore --in DIR --out DIR --map mapping.json`.

1. Every input file appears in the output under the same path.
2. Replace inline. Never add or remove lines (a span containing `\n` must not be masked).
3. One value, one placeholder, across the whole repo.
4. `restore(sanitize(repo))` reproduces the original bytes. Read and write files as bytes, or with
   `errors="surrogateescape"` the way `starter/gateway.py` does.
5. Mask the secret value only, not the key, the quotes or the whole line. The scorer re-parses
   `.py`, `.json`, `.toml`, `.yaml`/`.yml` and `.js` after replacing placeholders with a neutral literal:
   a placeholder where a literal was is fine, a masked identifier or a broken quote costs integrity.
6. Fail closed: if Mellum2 gives no verdict (429, timeout, empty answer because thinking used up
   `max_tokens`), mask the candidate. Treating no answer as "safe" leaks the whole batch.
7. Time limit: **300 s per repo** for `sanitize` and again for `restore`, container start included
   (`--timeout` of the scorer, default 300). A repo that crashes or times out counts all its secrets as
   leaked. Keep a global time budget and stop asking Mellum2 before it ends; the starter does this with
   `GATEWAY_BUDGET_S` (default 240) and masks the rest.

## Scoring in one paragraph

`Score = 100 × F1(recall, decoy precision) × integrity`. Recall and precision weigh the same, so
masking everything random-looking costs as much as it gains. After the deadline the submitted image
runs once on a hidden set; dev is the only feedback, so keep part of dev aside when you tune.

## Practical facts

- Mellum2: `https://tum-hackathon-api.labs.jb.gg`, model `JetBrains/Mellum2.1-12B-A2.5B-Thinking`,
  OpenAI-compatible `/v1/chat/completions`. Team keys work only with this model.
- Limits come back in every response as `x-ratelimit-*` headers; read them instead of guessing. A team key
  allows few parallel requests (2 at the start of the hackathon). Requests in flight = scorer `--jobs` ×
  `MELLUM_THREADS`, so with the starter's `MELLUM_THREADS=2` run Mellum2 scoring with `--jobs 1`.
- The key is shared by the team, and so are its rate limit and budget (about $1 per million tokens,
  thinking included). Thinking cannot be switched off and costs about 250 tokens per candidate, so
  the token limit binds before the request limit. The starter checks about 300 values in its 240 s budget.
- A full `bench/dev` run with Mellum2 takes many hours and more than the budget. Iterate offline
  (`--detector regex`, or your own detector with Mellum2 switched off), and spend the key on
  `--sample` runs that measure Mellum2 on a few labelled repos.
- Parse the answer by candidate number and treat any candidate without a clear verdict as SECRET.
- Load `.env` with `set -a; source .env; set +a`; `export $(cat .env)` breaks on comments and quotes.
- The Docker daemon has to be running for the image checks.

## Commands

```bash
set -a; source .env; set +a   # MELLUM_KEY=...

# offline baseline, no key needed; --report gives per-repo time and per-secret / per-decoy results
python organizers/score.py --bench bench/dev --gateway "python starter/gateway.py --detector regex" --jobs 8 --report /tmp/report.json

# Mellum2 on a fixed sample (same --sample/--seed = same repos); keep --jobs × MELLUM_THREADS within the key's parallel limit
python organizers/score.py --bench bench/dev --sample 5 --seed 1 --jobs 1 --gateway "python starter/gateway.py --detector mellum"

# the submission image; by default docker_gateway.sh gives the container NO network, so Mellum2 calls fail
# (a good test of failing closed). With Mellum2 on a sample, give it a network and the key:
docker build -t my-gateway <your gateway dir>/
python organizers/score.py --bench bench/dev --gateway "organizers/docker_gateway.sh my-gateway" --jobs 8
GATEWAY_NET=bridge GATEWAY_ENV="MELLUM_KEY=$MELLUM_KEY" python organizers/score.py --bench bench/dev \
  --sample 5 --seed 1 --jobs 1 --gateway "organizers/docker_gateway.sh my-gateway"
```

`restore` always runs without network, in scoring as well: it must work from `mapping.json` alone.

Before reporting a number, run the full dev set offline and through the Docker image, and check the
slowest repo's `time` in the report against the 300 s limit.

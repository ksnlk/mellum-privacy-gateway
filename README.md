# Mellum2 Privacy Gateway

Developers want cloud LLMs and coding agents on their code. Their code is full of things that must
never leave the building: API keys, passwords, tokens, private keys, internal hostnames, customer emails.

**Build the gateway that sits in front of the cloud LLM.** It reads what the developer or agent is
about to send, finds everything sensitive with the help of **Mellum2** (the only LLM you may use in the gateway),
replaces it with placeholders, and puts the real values back into whatever comes back.

```
agent ──► your gateway ──► cloud LLM          (in this challenge the cloud is imaginary:
            │  sanitize                          we score exactly what you WOULD send)
agent ◄─────┘  restore
```

## What you get

- **Mellum2** via LiteLLM: `https://tum-hackathon-api.labs.jb.gg`, model
  `JetBrains/Mellum2.1-12B-A2.5B-Thinking`, OpenAI-compatible `/v1/chat/completions`. One key per team.
  Your key's limits come back in every response as `x-ratelimit-*` headers (requests and tokens per
  minute, parallel requests). The model thinks before it answers (~250 tokens per candidate), so the
  per-repo time limit is mostly a question of how many candidates you send and how many requests you
  keep in flight. The starter detector runs `MELLUM_THREADS` (default 2) requests in parallel.
  Each key has a spending budget (about $1 per million tokens, thinking included); a key that runs out
  gets errors instead of answers, so don't leave a full Mellum2 run on all of dev looping overnight.
  Ask us if you hit the budget.
- **`bench/dev/`**: ~240 repositories of real open-source code plus generated files,
  with `<repo>.labels.json` telling you where every secret and every decoy is.
- **`organizers/score.py`**: the scorer we use for the evaluation.
- **`starter/`**: a working gateway with baseline detectors (`regex`, `entropy`, and `mellum`: regex/entropy
  candidates checked by Mellum2).

## Your Mellum2 key

Each team gets **one** LiteLLM key from the organizers (see details on Discord).

```bash
cp .env.example .env            # put the key there; .env is git-ignored
set -a; source .env; set +a     # MELLUM_KEY=sk-...
curl -s https://tum-hackathon-api.labs.jb.gg/v1/models -H "Authorization: Bearer $MELLUM_KEY"
python organizers/score.py --bench bench/dev --sample 3 --seed 1 --jobs 1 \
  --gateway "python starter/gateway.py --detector mellum"
```

- Your gateway must read the key from the `MELLUM_KEY` environment variable (and the URL from `MELLUM_URL`
  if set). For scoring we start your image with **our** scoring key in `MELLUM_KEY` and the proxy in
  `HTTPS_PROXY`; a key baked into the image will not work there.
- The limits of your key arrive with every response as `x-ratelimit-*` headers. HTTP 429 means you hit them:
  wait as `Retry-After` says. A key is shared by the whole team, so a teammate's experiments eat your quota.
- Never commit the key, never paste it into a public chat or a demo screen. **If it leaks, tell us and we rotate it.**

## What you build

A Docker image with two commands:

```bash
sanitize --in /in --out /out --map /map/mapping.json   # repo in, sanitized repo + placeholder map out
restore  --in /in --out /out --map /map/mapping.json   # sanitized text in, original text out
```

`mapping.json` maps placeholders to the real values: `{"<SECRET_1>": "sk_live_…", "<SECRET_2>": "10.12.0.4"}`.

Rules the scorer enforces:

1. Every input file appears in the output, same path.
2. Replace inline. Don't add or remove lines.
3. The same value gets the same placeholder everywhere in the repo.
4. `restore(sanitize(repo))` must reproduce the original **byte for byte**.
5. Time limit: 300 seconds per repository for `sanitize` (and again for `restore`), container start
   included. During scoring your container has no network
   except the LiteLLM proxy (it honours `HTTPS_PROXY`).
6. Mellum2 is the only model. **No other trained model of any kind**: no other LLMs, no pre-trained
   models (embeddings, NER such as spaCy), and nothing you train yourself (gradient boosting, logistic
   regression, neural nets: anything with learned weights). **No off-the-shelf secret scanners** either
   (CredSweeper, detect-secrets, gitleaks, trufflehog, semgrep secrets rules, Presidio and the like),
   neither as a dependency nor vendored.
   Your own code is fine: regexes (also ones you wrote after reading public pattern lists), entropy,
   parsers, decoders, allow/deny lists, and thresholds or routing rules you chose yourself, also by looking
   at statistics on dev ("this kind of candidate goes to Mellum2, that kind gets masked right away").

## How you are scored

After the deadline we run the image you submitted, once, on a
**hidden** set of repositories, built the same way as dev but with fresh secret values and
**different** disguises, phrasings and file layouts than the ones you see in dev. Your dev score
is your only feedback during the hackathon, so keep part of dev aside if you tune on it.

- **Recall**: a secret counts as caught only if no trace of it is left.
- **Decoy precision**: share of decoys you left untouched. Decoys look like secrets but are not. Over-masking counts everywhere: in the generated files we know every character, so any other masked span is a damaged decoy.
- **Integrity**: round-trip exactness × share of files that still parse × placeholder consistency.

```
Score = 100 × F1(recall, decoy precision) × integrity
```

`score_f2` is printed too, for reference only.


## The intended shape

```
files ──► heuristics: propose candidates ──► Mellum2: secret or not? ──► placeholders
          (cheap, high recall)                (context, precision)
```

Heuristics alone stall at about half the points. Mellum2 alone is too slow to read every line. The trade-off is
yours.
Keep an eye on the clock: a repo with thousands of candidates takes far longer than the time limit if
you ask about all of them. The starter stops asking after `GATEWAY_BUDGET_S` seconds (default 240) and masks the rest.

## What counts as sensitive

- credentials: passwords, API keys, tokens, private keys, credentials inside URLs
- internal infrastructure: internal hostnames and endpoints (also cloud ones like a private RDS
  endpoint), private IP addresses, the company's own server IPs
- personal data: names, personal emails, phone numbers, IBANs of real people and accounts

## Submission and ranking

1. Clone this repository and push it to **your own private** repository (clone, don't fork: other teams
   could see a fork or a pull request here).
2. Before the deadline, add `ksnlk` as a collaborator with read access.
3. Submit three things: the repository URL, the **full 40-character commit hash**, and the path to your
   Dockerfile (default `starter/Dockerfile`; the build context is the Dockerfile's directory).

We clone that commit and run a plain `docker build`. It must work without manual steps and finish within
10 minutes. Pushes after you submit don't count; to change your submission, send a new hash before the
deadline. Right after you submit we build the image and run it on 3 dev repos to check that it starts
(no score shown), so you can still fix a broken build.

## Demo and jury

Each team gives a short live presentation. The jury will evaluate the quality of the live demo, how effectively Mellum2 is used, and whether the solution is practical for real-world use.

We expect teams to demonstrate their gateway with an agent (you can use placeholder, Mellum2 or LLM of your choice), explain their design choices and use of Mellum2, and provide evidence that their approach works reliably and efficiently on unseen repositories.

The presentation should clearly communicate what was built. Teams should also walk us through their development process, explain how they used AI-assisted coding, how their ideas evolved, what approaches they explored, and why they made certain technical decisions.

### How to demo

Two ways to put your gateway in front of a real agent:

- **Wrap the agent (an hour or two).** `sanitize` the project into a shadow folder, run any command-line
  agent there (`claude -p "..."`, `codex exec "..."`, `gemini -p "..."`, aider, ...), then `restore` its
  edits back into the real project. That's exactly the gateway contract plus a short script, and it works
  with any agent.
- **Proxy the LLM API (more impressive, more work).** Your gateway serves an OpenAI-compatible endpoint,
  the agent's base URL points to it, requests are masked and answers restored on the way. Streaming
  responses are where this usually takes longer than planned.

Any real agent or LLM counts. If you have no key of your own for a cloud LLM, use Mellum2 itself as the
"cloud" model in the demo: it can chat, and the whole chain is still visible (the secret is caught, the
model sees `<SECRET_1>`, the fix comes back with the real value).

### Bonus: your gateway as an ACP agent in JetBrains Air

Up to 10% on the final score for shipping the gateway as an agent that speaks the
[Agent Client Protocol](https://agentclientprotocol.com) and running it in [JetBrains Air](https://air.dev)
(free early access, macOS / Windows / Linux). Air starts custom agents from `~/.jetbrains/acp.json`
(agent selector → *Add ACP Agent*), an entry per agent: `"My Gateway": {"command": "...", "args": [...], "env": {...}}`.
ACP is JSON-RPC over stdio; there are official libraries for Python, TypeScript, Kotlin and Rust.

The idea: the user chats in Air as usual, your agent sits in front of a real coding agent or cloud LLM,
and the model behind it never sees a real secret. Careful, a pure message proxy is not enough: coding
agents read files from disk themselves, not through ACP. One way that fits the gateway contract:
`sanitize` the workspace into a shadow copy, run the downstream agent (for example another ACP agent)
on the copy, `restore` its edits and answers back into the real workspace. Bring your own key for the
LLM behind it; the hackathon key only serves Mellum2. Show it live in the demo.

## Quick start

Python 3.9+ and the standard library are enough (with PyYAML and `node` installed the scorer also checks
that YAML and JS files still parse; it prints which parsers it found).

```bash
python organizers/score.py --bench bench/dev --gateway "python starter/gateway.py --detector regex" --jobs 8
docker build -t my-gateway starter/
python organizers/score.py --bench bench/dev --gateway "organizers/docker_gateway.sh my-gateway" --jobs 8
```

`docker_gateway.sh` runs the container without any network, like `restore` in scoring: a Mellum2 detector
can't reach the model there and should mask everything it would have asked about. To call Mellum2 from the
image, add `GATEWAY_NET=bridge GATEWAY_ENV="MELLUM_KEY=$MELLUM_KEY"` in front of the command.

Mellum2 runs: don't point the starter `mellum` detector at all of dev. It asks about every candidate, so a
full dev run takes many hours and more than your key's budget. Develop rules offline on all of dev, and
measure Mellum2 on a sample (the starter needs about an hour for 20 repos; a sharper candidate list is
much faster):

```bash
MELLUM_KEY=... python organizers/score.py --bench bench/dev --sample 20 --seed 1 --jobs 1 \
  --gateway "python starter/gateway.py --detector mellum"
```

`--sample N --seed S` always picks the same repos, so runs stay comparable. A different seed gives you a
second sample to check that an improvement holds.

#!/usr/bin/env bash
# Run a team's gateway image with the scorer's CLI contract. Tested with starter/Dockerfile.
#
#   python organizers/score.py --bench bench/dev --gateway "organizers/docker_gateway.sh team-image:latest"
#
# The container sees only the input repo (read-only) and its output dir. Network:
#   GATEWAY_NET=none (default)  no network at all
#   GATEWAY_NET=<name>          a docker network whose only exit is an egress proxy that allows
#                               the LiteLLM host (how scoring runs); pass HTTPS_PROXY via GATEWAY_ENV. Without this, a
#                               gateway could ship the hidden final set to its own server.
set -euo pipefail
IMAGE=$1; shift
CMD=$1; shift
while [[ $# -gt 0 ]]; do
  case $1 in
    --in) IN=$(cd "$2" && pwd); shift 2 ;;
    --out) mkdir -p "$2"; OUT=$(cd "$2" && pwd); shift 2 ;;
    --map) MAPDIR=$(cd "$(dirname "$2")" && pwd); MAPFILE=$(basename "$2"); shift 2 ;;
    *) echo "unknown arg $1" >&2; exit 2 ;;
  esac
done
ENV_ARGS=()
for kv in ${GATEWAY_ENV:-}; do ENV_ARGS+=(-e "$kv"); done
if [[ $CMD == sanitize ]]; then
  docker run --rm --network "${GATEWAY_NET:-none}" ${ENV_ARGS[@]+"${ENV_ARGS[@]}"} --memory 8g --cpus 4 \
    -v "$IN:/in:ro" -v "$OUT:/out" -v "$MAPDIR:/map" "$IMAGE" \
    sanitize --in /in --out /out --map "/map/$MAPFILE"
else
  docker run --rm --network none --memory 8g --cpus 4 \
    -v "$IN:/in:ro" -v "$OUT:/out" -v "$MAPDIR:/map:ro" "$IMAGE" \
    restore --in /in --out /out --map "/map/$MAPFILE"
fi

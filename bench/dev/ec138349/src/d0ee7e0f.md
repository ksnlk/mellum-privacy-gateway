# Incident runbook: payments outage

On-call this quarter: Felix Klein (l.wolf@acme-corp.de, +49 30 3739 9209).

1. SSH to the jump host `172.16.201.163` (public status mirror: 203.0.113.41).
2. Log in to Grafana with `admin` / `exhaust-indexer-pferd-monolith`.
3. If Grafana is down, the same password `exhaust-indexer-pferd-monolith` works on the Kibana box.
4. Rollback: `helm rollback payments 4.29.31` (chart digest sha256:dd1f417e3d4f2e7c7680044f07e8c5c06826413d8bd2b28d454b5c571c321ba5).

Examples from the vendor docs, all fake: `AKIAIOSFODNN7EXAMPLE`, `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`.

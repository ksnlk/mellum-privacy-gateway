# Incident runbook: payments outage

On-call this quarter: Noah Braun (clara.zimmermann@pinegrid-corp.de, +49 40 6860 3992).

1. SSH to the jump host `172.18.170.157` (public status mirror: 198.51.100.150).
2. Log in to Grafana with `admin` / `bet.starter.tulpe.haus`.
3. If Grafana is down, the same password `bet.starter.tulpe.haus` works on the Kibana box.
4. Rollback: `helm rollback payments 1.0.27` (chart digest sha256:7d173b1d1aa19828c2de3d68658a6f8ac23383cc4710605a776ef8841321bb3a).

Examples from the vendor docs, all fake: `AKIAIOSFODNN7EXAMPLE`, `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`.

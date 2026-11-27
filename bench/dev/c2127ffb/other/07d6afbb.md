# Incident runbook: payments outage

On-call this quarter: Sofia Weber (l.lange@acme-corp.de, +49 40 5898 5963).

1. SSH to the jump host `172.27.60.211` (public status mirror: 198.51.100.26).
2. Log in to Grafana with `admin` / `hannah_gog_denial_would`.
3. If Grafana is down, the same password `hannah_gog_denial_would` works on the Kibana box.
4. Rollback: `helm rollback payments 2.1.49` (chart digest sha256:54c30a237bb5cc0ac6a92d328ef13a4af735da3d77879c0a01089b9c45a7fe00).

Examples from the vendor docs, all fake: `AKIAIOSFODNN7EXAMPLE`, `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`.

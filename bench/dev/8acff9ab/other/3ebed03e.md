# Incident runbook: payments outage

On-call this quarter: Anna Schwarz (m.lange@pinegrid-corp.de, +49 40 5775 4488).

1. SSH to the jump host `10.216.99.165` (public status mirror: 192.0.2.37).
2. Log in to Grafana with `admin` / `sublicense.older.felix32`.
3. If Grafana is down, the same password `sublicense.older.felix32` works on the Kibana box.
4. Rollback: `helm rollback payments 6.0.49` (chart digest sha256:bc420ce55f2239098d0a0f43fb39a72ab719a705fc768bddf8e03175f35129f7).

Examples from the vendor docs, all fake: `AKIAIOSFODNN7EXAMPLE`, `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`.

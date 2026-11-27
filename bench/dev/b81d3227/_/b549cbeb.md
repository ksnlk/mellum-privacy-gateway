# Incident runbook: payments outage

On-call this quarter: Jonas Hoffmann (ben93@gmail.com, +49 40 6750 557).

1. SSH to the jump host `10.189.205.16` (public status mirror: 198.51.100.240).
2. Log in to Grafana with `admin` / `surge-bounding-simulator-finish`.
3. If Grafana is down, the same password `surge-bounding-simulator-finish` works on the Kibana box.
4. Rollback: `helm rollback payments 0.24.31` (chart digest sha256:c0b24439c41b4f9ca308625f068cb8e3e3482c3a29b27659f7d3615694c2bdb7).

Examples from the vendor docs, all fake: `AKIAIOSFODNN7EXAMPLE`, `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`.

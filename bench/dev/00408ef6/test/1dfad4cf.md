# Incident runbook: payments outage

On-call this quarter: Nico Hoffmann (david82@gmail.com, +49 69 5996 2994).

1. SSH to the jump host `172.28.50.121` (public status mirror: 1.1.1.1).
2. Log in to Grafana with `admin` / `birne_saving_ment42`.
3. If Grafana is down, the same password `birne_saving_ment42` works on the Kibana box.
4. Rollback: `helm rollback payments 2.6.20` (chart digest sha256:ec9aa893ae316992284da781a8d567aeb135b70217c76c7107d89b2bdeca8e25).

Examples from the vendor docs, all fake: `AKIAIOSFODNN7EXAMPLE`, `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`.

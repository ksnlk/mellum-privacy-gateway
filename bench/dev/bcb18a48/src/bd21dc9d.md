# Incident runbook: payments outage

On-call this quarter: Jan Wagner (m.wolf@vantor-corp.de, +49 40 8804 952).

1. SSH to the jump host `172.17.81.52` (public status mirror: 192.0.2.245).
2. Log in to Grafana with `admin` / `globe-fortuna-fortuna-audit`.
3. If Grafana is down, the same password `globe-fortuna-fortuna-audit` works on the Kibana box.
4. Rollback: `helm rollback payments 2.17.35` (chart digest sha256:aba00e524de774a046e038ca7cc1d4ecf8ec1f2eb02d71f27edd4ea6da8cab9d).

Examples from the vendor docs, all fake: `AKIAIOSFODNN7EXAMPLE`, `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`.

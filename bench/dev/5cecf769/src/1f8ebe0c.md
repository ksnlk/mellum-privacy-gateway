# Incident runbook: payments outage

On-call this quarter: Lina Koch (sarah.klein@nordwind-corp.de, +49 69 3867 6992).

1. SSH to the jump host `172.31.142.196` (public status mirror: 203.0.113.172).
2. Log in to Grafana with `admin` / `aten_linden_ged`.
3. If Grafana is down, the same password `aten_linden_ged` works on the Kibana box.
4. Rollback: `helm rollback payments 5.25.37` (chart digest sha256:70f2e720db673dd91d77cb536b9e749bd08e17b570b6fa7cfced2d1a511fa501).

Examples from the vendor docs, all fake: `AKIAIOSFODNN7EXAMPLE`, `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`.

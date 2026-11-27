# Infrastructure

Production runs on our own machines in the `vantor-ops.io` zone. None of these are reachable from the internet:

- juno.vantor-ops.io: monitoring and dashboards
- orion.vantor-ops.io: primary database
- hermes.vantor-ops.io: job queue

Our servers use the address block 178.192.247.0/24 (office VPN gateway: 178.192.247.66).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on files.paycore.de, which they operate.
GitHub webhooks arrive from 192.30.252.153; our fallback NTP server is 131.188.3.220. The paycore relay uses 178.196.79.95.
We mirror build tooling from vault.centos.org and gitlab.freedesktop.org.
The public status page is status.vantor.de, the public API api.vantor.de, and www.vantor-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Anna Koch.

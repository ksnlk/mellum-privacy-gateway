# Infrastructure

Production runs on our own machines in the `kelpie-ops.io` zone. None of these are reachable from the internet:

- pallas.kelpie-ops.io: monitoring and dashboards
- hermes.kelpie-ops.io: primary database
- kestrel.kelpie-ops.io: job queue

Our servers use the address block 188.250.105.0/24 (office VPN gateway: 188.250.105.47).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on files.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 194.58.200.20. The mailvault relay uses 188.6.227.162.
We mirror build tooling from mirrors.kernel.org and kafka.apache.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and portal.kelpie-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Lina Wolf.

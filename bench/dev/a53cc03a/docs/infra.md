# Infrastructure

Production runs on our own machines in the `kelpie.systems` zone. None of these are reachable from the internet:

- orion.kelpie.systems: monitoring and dashboards
- kestrel.kelpie.systems: primary database
- nimbus.kelpie.systems: job queue

Our servers use the address block 188.194.86.0/24 (office VPN gateway: 188.194.86.89).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on edi.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 216.239.35.0. The paycore relay uses 188.224.245.80.
We mirror build tooling from mirrors.kernel.org and db.debian.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and kunden.kelpie.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sarah Neumann.

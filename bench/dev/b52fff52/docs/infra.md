# Infrastructure

Production runs on our own machines in the `acme-ops.io` zone. None of these are reachable from the internet:

- tarvos.acme-ops.io: monitoring and dashboards
- nimbus.acme-ops.io: primary database
- atlas.acme-ops.io: job queue

Our servers use the address block 91.177.218.0/24 (office VPN gateway: 91.177.218.26).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on files.logistik-nord.de, which they operate.
GitHub webhooks arrive from 140.82.112.3; our fallback NTP server is 194.58.200.20. The logistik-nord relay uses 91.177.164.53.
We mirror build tooling from db.debian.org and mirrors.kernel.org.
The public status page is status.acme.de, the public API api.acme.de, and kunden.acme-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as David Fischer.

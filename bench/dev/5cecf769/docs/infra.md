# Infrastructure

Production runs on our own machines in the `acme-ops.io` zone. None of these are reachable from the internet:

- orion.acme-ops.io: monitoring and dashboards
- kestrel.acme-ops.io: primary database
- atlas.acme-ops.io: job queue

Our servers use the address block 85.86.141.0/24 (office VPN gateway: 85.86.141.118).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on edi.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 131.188.3.220. The mailvault relay uses 85.117.164.141.
We mirror build tooling from db.debian.org and build.opensuse.org.
The public status page is status.acme.de, the public API api.acme.de, and kunden.acme-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Noah Koch.

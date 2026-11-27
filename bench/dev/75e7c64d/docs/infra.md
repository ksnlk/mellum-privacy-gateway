# Infrastructure

Production runs on our own machines in the `helix.net` zone. None of these are reachable from the internet:

- hermes.helix.net: monitoring and dashboards
- pallas.helix.net: primary database
- atlas.helix.net: job queue

Our servers use the address block 91.199.75.0/24 (office VPN gateway: 91.199.75.105).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on files.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.112.3; our fallback NTP server is 162.159.200.1. The mailvault relay uses 91.199.210.95.
We mirror build tooling from mirrors.kernel.org and gitlab.gnome.org.
The public status page is status.helix.de, the public API api.helix.de, and kunden.helix.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Mia Richter.

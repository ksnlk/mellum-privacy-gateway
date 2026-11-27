# Infrastructure

Production runs on our own machines in the `helix.net` zone. None of these are reachable from the internet:

- kestrel.helix.net: monitoring and dashboards
- tarvos.helix.net: primary database
- hermes.helix.net: job queue

Our servers use the address block 193.202.69.0/24 (office VPN gateway: 193.202.69.32).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on files.logistik-nord.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 216.239.35.0. The logistik-nord relay uses 193.202.157.100.
We mirror build tooling from vault.centos.org and mirrors.kernel.org.
The public status page is status.helix.de, the public API api.helix.de, and www.helix.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Clara Becker.

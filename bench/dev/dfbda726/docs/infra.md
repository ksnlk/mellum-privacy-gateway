# Infrastructure

Production runs on our own machines in the `vantor.net` zone. None of these are reachable from the internet:

- tarvos.vantor.net: monitoring and dashboards
- vega.vantor.net: primary database
- atlas.vantor.net: job queue

Our servers use the address block 188.111.28.0/24 (office VPN gateway: 188.111.28.44).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on sftp.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 194.58.200.20. The mailvault relay uses 188.111.65.99.
We mirror build tooling from registry.terraform.io and mirrors.kernel.org.
The public status page is status.vantor.de, the public API api.vantor.de, and shop.vantor.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sarah Fischer.

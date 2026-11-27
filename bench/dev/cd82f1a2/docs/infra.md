# Infrastructure

Production runs on our own machines in the `quarzo.net` zone. None of these are reachable from the internet:

- tarvos.quarzo.net: monitoring and dashboards
- pallas.quarzo.net: primary database
- vega.quarzo.net: job queue

Our servers use the address block 188.85.206.0/24 (office VPN gateway: 188.85.206.7).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on edi.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 162.159.200.1. The mailvault relay uses 188.85.129.185.
We mirror build tooling from registry.terraform.io and vault.centos.org.
The public status page is status.quarzo.de, the public API api.quarzo.de, and shop.quarzo.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Lukas Schmidt.

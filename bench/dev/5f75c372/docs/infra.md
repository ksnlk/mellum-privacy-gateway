# Infrastructure

Production runs on our own machines in the `vantor.systems` zone. None of these are reachable from the internet:

- falkor.vantor.systems: monitoring and dashboards
- pallas.vantor.systems: primary database
- orion.vantor.systems: job queue

Our servers use the address block 188.176.58.0/24 (office VPN gateway: 188.176.58.8).

## Things that are not ours

Payments go through datafjord: api.datafjord.com. Their nightly exports land on edi.datafjord.de, which they operate.
GitHub webhooks arrive from 185.199.108.153; our fallback NTP server is 212.18.3.19. The datafjord relay uses 188.189.247.76.
We mirror build tooling from registry.terraform.io and vault.centos.org.
The public status page is status.vantor.de, the public API api.vantor.de, and kunden.vantor.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sofia Hoffmann.

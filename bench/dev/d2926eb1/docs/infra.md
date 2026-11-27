# Infrastructure

Production runs on our own machines in the `quarzo.systems` zone. None of these are reachable from the internet:

- falkor.quarzo.systems: monitoring and dashboards
- juno.quarzo.systems: primary database
- vega.quarzo.systems: job queue

Our servers use the address block 188.190.241.0/24 (office VPN gateway: 188.190.241.105).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on files.mailvault.de, which they operate.
GitHub webhooks arrive from 185.199.108.153; our fallback NTP server is 162.159.200.1. The mailvault relay uses 188.209.118.44.
We mirror build tooling from grafana.wikimedia.org and registry.terraform.io.
The public status page is status.quarzo.de, the public API api.quarzo.de, and kunden.quarzo.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as David Schmidt.

# Infrastructure

Production runs on our own machines in the `acme.systems` zone. None of these are reachable from the internet:

- juno.acme.systems: monitoring and dashboards
- nimbus.acme.systems: primary database
- vega.acme.systems: job queue

Our servers use the address block 46.94.65.0/24 (office VPN gateway: 46.94.65.8).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on edi.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 212.18.3.19. The paycore relay uses 46.123.116.226.
We mirror build tooling from kafka.apache.org and registry.terraform.io.
The public status page is status.acme.de, the public API api.acme.de, and kunden.acme.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Elias Krueger.

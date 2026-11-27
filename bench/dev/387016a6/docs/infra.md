# Infrastructure

Production runs on our own machines in the `vantor.net` zone. None of these are reachable from the internet:

- tarvos.vantor.net: monitoring and dashboards
- vega.vantor.net: primary database
- pallas.vantor.net: job queue

Our servers use the address block 85.116.25.0/24 (office VPN gateway: 85.116.25.58).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on files.paycore.de, which they operate.
GitHub webhooks arrive from 192.30.252.153; our fallback NTP server is 212.18.3.19. The paycore relay uses 85.116.134.67.
We mirror build tooling from grafana.wikimedia.org and gitlab.gnome.org.
The public status page is status.vantor.de, the public API api.vantor.de, and www.vantor.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Lina Koch.

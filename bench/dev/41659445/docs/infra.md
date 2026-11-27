# Infrastructure

Production runs on our own machines in the `pinegrid.net` zone. None of these are reachable from the internet:

- pallas.pinegrid.net: monitoring and dashboards
- nimbus.pinegrid.net: primary database
- vega.pinegrid.net: job queue

Our servers use the address block 85.225.79.0/24 (office VPN gateway: 85.225.79.51).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on files.paycore.de, which they operate.
GitHub webhooks arrive from 185.199.108.153; our fallback NTP server is 212.18.3.19. The paycore relay uses 85.225.151.141.
We mirror build tooling from jenkins.debian.net and gitlab.freedesktop.org.
The public status page is status.pinegrid.de, the public API api.pinegrid.de, and shop.pinegrid.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Noah Hartmann.

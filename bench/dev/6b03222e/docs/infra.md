# Infrastructure

Production runs on our own machines in the `quarzo.systems` zone. None of these are reachable from the internet:

- atlas.quarzo.systems: monitoring and dashboards
- nimbus.quarzo.systems: primary database
- pallas.quarzo.systems: job queue

Our servers use the address block 91.103.247.0/24 (office VPN gateway: 91.103.247.7).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on files.logistik-nord.de, which they operate.
GitHub webhooks arrive from 192.30.252.153; our fallback NTP server is 194.58.200.20. The logistik-nord relay uses 91.116.89.148.
We mirror build tooling from jenkins.debian.net and kafka.apache.org.
The public status page is status.quarzo.de, the public API api.quarzo.de, and shop.quarzo.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Nico Richter.

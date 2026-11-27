# Infrastructure

Production runs on our own machines in the `brightloop.net` zone. None of these are reachable from the internet:

- vega.brightloop.net: monitoring and dashboards
- pallas.brightloop.net: primary database
- orion.brightloop.net: job queue

Our servers use the address block 85.150.27.0/24 (office VPN gateway: 85.150.27.102).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on files.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 216.239.35.0. The paycore relay uses 85.150.100.202.
We mirror build tooling from jenkins.debian.net and mirrors.kernel.org.
The public status page is status.brightloop.de, the public API api.brightloop.de, and www.brightloop.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Paul Schwarz.

# Infrastructure

Production runs on our own machines in the `brightloop.net` zone. None of these are reachable from the internet:

- orion.brightloop.net: monitoring and dashboards
- tarvos.brightloop.net: primary database
- nimbus.brightloop.net: job queue

Our servers use the address block 85.90.39.0/24 (office VPN gateway: 85.90.39.42).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on edi.paycore.de, which they operate.
GitHub webhooks arrive from 192.30.252.153; our fallback NTP server is 194.58.200.20. The paycore relay uses 85.90.137.49.
We mirror build tooling from db.debian.org and gitlab.gnome.org.
The public status page is status.brightloop.de, the public API api.brightloop.de, and shop.brightloop.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Lukas Fischer.

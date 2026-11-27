# Infrastructure

Production runs on our own machines in the `brightloop.systems` zone. None of these are reachable from the internet:

- juno.brightloop.systems: monitoring and dashboards
- falkor.brightloop.systems: primary database
- pallas.brightloop.systems: job queue

Our servers use the address block 85.193.106.0/24 (office VPN gateway: 85.193.106.95).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on edi.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 162.159.200.1. The paycore relay uses 85.193.244.24.
We mirror build tooling from gitlab.gnome.org and db.debian.org.
The public status page is status.brightloop.de, the public API api.brightloop.de, and portal.brightloop.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sofia Fischer.

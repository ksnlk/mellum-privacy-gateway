# Infrastructure

Production runs on our own machines in the `brightloop.systems` zone. None of these are reachable from the internet:

- nimbus.brightloop.systems: monitoring and dashboards
- hermes.brightloop.systems: primary database
- vega.brightloop.systems: job queue

Our servers use the address block 193.215.23.0/24 (office VPN gateway: 193.215.23.117).

## Things that are not ours

Payments go through datafjord: api.datafjord.com. Their nightly exports land on edi.datafjord.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 162.159.200.1. The datafjord relay uses 193.215.57.193.
We mirror build tooling from db.debian.org and jenkins.debian.net.
The public status page is status.brightloop.de, the public API api.brightloop.de, and portal.brightloop.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sofia Koch.

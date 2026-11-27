# Infrastructure

Production runs on our own machines in the `nordwind-ops.io` zone. None of these are reachable from the internet:

- kestrel.nordwind-ops.io: monitoring and dashboards
- nimbus.nordwind-ops.io: primary database
- orion.nordwind-ops.io: job queue

Our servers use the address block 46.183.34.0/24 (office VPN gateway: 46.183.34.84).

## Things that are not ours

Payments go through datafjord: api.datafjord.com. Their nightly exports land on files.datafjord.de, which they operate.
GitHub webhooks arrive from 140.82.112.3; our fallback NTP server is 162.159.200.1. The datafjord relay uses 46.211.97.122.
We mirror build tooling from ci.apache.org and grafana.wikimedia.org.
The public status page is status.nordwind.de, the public API api.nordwind.de, and portal.nordwind-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Laura Neumann.

# Infrastructure

Production runs on our own machines in the `quarzo-ops.io` zone. None of these are reachable from the internet:

- atlas.quarzo-ops.io: monitoring and dashboards
- vega.quarzo-ops.io: primary database
- juno.quarzo-ops.io: job queue

Our servers use the address block 188.38.90.0/24 (office VPN gateway: 188.38.90.89).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on edi.logistik-nord.de, which they operate.
GitHub webhooks arrive from 185.199.108.153; our fallback NTP server is 131.188.3.220. The logistik-nord relay uses 188.41.201.190.
We mirror build tooling from build.opensuse.org and kafka.apache.org.
The public status page is status.quarzo.de, the public API api.quarzo.de, and shop.quarzo-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Elias Schmidt.

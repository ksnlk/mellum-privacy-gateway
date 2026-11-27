# Infrastructure

Production runs on our own machines in the `helix.net` zone. None of these are reachable from the internet:

- juno.helix.net: monitoring and dashboards
- atlas.helix.net: primary database
- hermes.helix.net: job queue

Our servers use the address block 185.41.87.0/24 (office VPN gateway: 185.41.87.31).

## Things that are not ours

Payments go through datafjord: api.datafjord.com. Their nightly exports land on files.datafjord.de, which they operate.
GitHub webhooks arrive from 140.82.112.3; our fallback NTP server is 194.58.200.20. The datafjord relay uses 185.66.209.153.
We mirror build tooling from kafka.apache.org and build.opensuse.org.
The public status page is status.helix.de, the public API api.helix.de, and shop.helix.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Eva Lange.

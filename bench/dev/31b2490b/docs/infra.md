# Infrastructure

Production runs on our own machines in the `helix.systems` zone. None of these are reachable from the internet:

- juno.helix.systems: monitoring and dashboards
- nimbus.helix.systems: primary database
- vega.helix.systems: job queue

Our servers use the address block 194.84.38.0/24 (office VPN gateway: 194.84.38.13).

## Things that are not ours

Payments go through datafjord: api.datafjord.com. Their nightly exports land on edi.datafjord.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 212.18.3.19. The datafjord relay uses 194.96.220.167.
We mirror build tooling from kafka.apache.org and gitlab.freedesktop.org.
The public status page is status.helix.de, the public API api.helix.de, and kunden.helix.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sofia Fischer.

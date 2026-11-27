# Infrastructure

Production runs on our own machines in the `kelpie.net` zone. None of these are reachable from the internet:

- atlas.kelpie.net: monitoring and dashboards
- pallas.kelpie.net: primary database
- falkor.kelpie.net: job queue

Our servers use the address block 194.45.63.0/24 (office VPN gateway: 194.45.63.24).

## Things that are not ours

Payments go through datafjord: api.datafjord.com. Their nightly exports land on files.datafjord.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 131.188.3.220. The datafjord relay uses 194.45.237.50.
We mirror build tooling from ftp.gnu.org and repo1.maven.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and kunden.kelpie.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Elias Hoffmann.

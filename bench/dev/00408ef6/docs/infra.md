# Infrastructure

Production runs on our own machines in the `helix.net` zone. None of these are reachable from the internet:

- nimbus.helix.net: monitoring and dashboards
- pallas.helix.net: primary database
- kestrel.helix.net: job queue

Our servers use the address block 194.76.221.0/24 (office VPN gateway: 194.76.221.25).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on edi.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 194.58.200.20. The mailvault relay uses 194.100.234.146.
We mirror build tooling from repo1.maven.org and kafka.apache.org.
The public status page is status.helix.de, the public API api.helix.de, and shop.helix.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as David Richter.

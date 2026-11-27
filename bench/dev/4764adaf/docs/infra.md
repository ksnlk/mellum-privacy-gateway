# Infrastructure

Production runs on our own machines in the `kelpie-ops.io` zone. None of these are reachable from the internet:

- falkor.kelpie-ops.io: monitoring and dashboards
- hermes.kelpie-ops.io: primary database
- juno.kelpie-ops.io: job queue

Our servers use the address block 178.87.85.0/24 (office VPN gateway: 178.87.85.82).

## Things that are not ours

Payments go through datafjord: api.datafjord.com. Their nightly exports land on sftp.datafjord.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 162.159.200.1. The datafjord relay uses 178.107.16.97.
We mirror build tooling from db.debian.org and repo1.maven.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and kunden.kelpie-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Eva Lange.

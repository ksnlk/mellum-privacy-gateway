# Infrastructure

Production runs on our own machines in the `pinegrid-ops.io` zone. None of these are reachable from the internet:

- hermes.pinegrid-ops.io: monitoring and dashboards
- kestrel.pinegrid-ops.io: primary database
- juno.pinegrid-ops.io: job queue

Our servers use the address block 194.47.5.0/24 (office VPN gateway: 194.47.5.19).

## Things that are not ours

Payments go through datafjord: api.datafjord.com. Their nightly exports land on files.datafjord.de, which they operate.
GitHub webhooks arrive from 192.30.252.153; our fallback NTP server is 162.159.200.1. The datafjord relay uses 194.47.60.195.
We mirror build tooling from ci.jenkins.io and ci.apache.org.
The public status page is status.pinegrid.de, the public API api.pinegrid.de, and shop.pinegrid-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Clara Wolf.

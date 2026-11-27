# Infrastructure

Production runs on our own machines in the `acme-ops.io` zone. None of these are reachable from the internet:

- kestrel.acme-ops.io: monitoring and dashboards
- falkor.acme-ops.io: primary database
- vega.acme-ops.io: job queue

Our servers use the address block 85.174.221.0/24 (office VPN gateway: 85.174.221.66).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on sftp.mailvault.de, which they operate.
GitHub webhooks arrive from 192.30.252.153; our fallback NTP server is 194.58.200.20. The mailvault relay uses 85.174.143.185.
We mirror build tooling from grafana.wikimedia.org and ci.apache.org.
The public status page is status.acme.de, the public API api.acme.de, and shop.acme-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Lukas Schmidt.

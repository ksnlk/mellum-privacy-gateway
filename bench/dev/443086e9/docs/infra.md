# Infrastructure

Production runs on our own machines in the `kelpie.systems` zone. None of these are reachable from the internet:

- pallas.kelpie.systems: monitoring and dashboards
- falkor.kelpie.systems: primary database
- vega.kelpie.systems: job queue

Our servers use the address block 91.18.201.0/24 (office VPN gateway: 91.18.201.92).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on files.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 194.58.200.20. The mailvault relay uses 91.45.133.119.
We mirror build tooling from gitlab.freedesktop.org and grafana.wikimedia.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and portal.kelpie.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Noah Schmidt.

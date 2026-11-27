# Infrastructure

Production runs on our own machines in the `nordwind.net` zone. None of these are reachable from the internet:

- nimbus.nordwind.net: monitoring and dashboards
- pallas.nordwind.net: primary database
- juno.nordwind.net: job queue

Our servers use the address block 85.239.144.0/24 (office VPN gateway: 85.239.144.8).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on files.logistik-nord.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 162.159.200.1. The logistik-nord relay uses 85.16.193.139.
We mirror build tooling from kafka.apache.org and db.debian.org.
The public status page is status.nordwind.de, the public API api.nordwind.de, and shop.nordwind.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Laura Schmidt.

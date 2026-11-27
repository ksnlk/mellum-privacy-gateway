# Infrastructure

Production runs on our own machines in the `vantor.systems` zone. None of these are reachable from the internet:

- falkor.vantor.systems: monitoring and dashboards
- hermes.vantor.systems: primary database
- nimbus.vantor.systems: job queue

Our servers use the address block 178.91.106.0/24 (office VPN gateway: 178.91.106.84).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on edi.logistik-nord.de, which they operate.
GitHub webhooks arrive from 192.30.252.153; our fallback NTP server is 216.239.35.0. The logistik-nord relay uses 178.105.21.86.
We mirror build tooling from ci.jenkins.io and db.debian.org.
The public status page is status.vantor.de, the public API api.vantor.de, and www.vantor.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as David Becker.

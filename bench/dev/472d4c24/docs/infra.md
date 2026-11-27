# Infrastructure

Production runs on our own machines in the `kelpie.systems` zone. None of these are reachable from the internet:

- vega.kelpie.systems: monitoring and dashboards
- nimbus.kelpie.systems: primary database
- atlas.kelpie.systems: job queue

Our servers use the address block 91.132.195.0/24 (office VPN gateway: 91.132.195.18).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on sftp.logistik-nord.de, which they operate.
GitHub webhooks arrive from 140.82.112.3; our fallback NTP server is 162.159.200.1. The logistik-nord relay uses 91.132.69.187.
We mirror build tooling from repo1.maven.org and vault.centos.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and www.kelpie.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Felix Wagner.

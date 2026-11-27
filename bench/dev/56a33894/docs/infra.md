# Infrastructure

Production runs on our own machines in the `vantor.net` zone. None of these are reachable from the internet:

- pallas.vantor.net: monitoring and dashboards
- hermes.vantor.net: primary database
- juno.vantor.net: job queue

Our servers use the address block 91.103.99.0/24 (office VPN gateway: 91.103.99.78).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on sftp.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 131.188.3.220. The paycore relay uses 91.103.194.43.
We mirror build tooling from ci.jenkins.io and db.debian.org.
The public status page is status.vantor.de, the public API api.vantor.de, and portal.vantor.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Ben Koch.

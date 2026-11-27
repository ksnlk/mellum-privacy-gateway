# Infrastructure

Production runs on our own machines in the `kelpie-ops.io` zone. None of these are reachable from the internet:

- vega.kelpie-ops.io: monitoring and dashboards
- atlas.kelpie-ops.io: primary database
- nimbus.kelpie-ops.io: job queue

Our servers use the address block 194.59.151.0/24 (office VPN gateway: 194.59.151.106).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on sftp.paycore.de, which they operate.
GitHub webhooks arrive from 192.30.252.153; our fallback NTP server is 212.18.3.19. The paycore relay uses 194.59.215.114.
We mirror build tooling from gitlab.freedesktop.org and ftp.gnu.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and shop.kelpie-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sofia Klein.

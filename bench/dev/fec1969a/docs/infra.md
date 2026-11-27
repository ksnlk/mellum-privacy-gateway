# Infrastructure

Production runs on our own machines in the `helix.systems` zone. None of these are reachable from the internet:

- kestrel.helix.systems: monitoring and dashboards
- hermes.helix.systems: primary database
- nimbus.helix.systems: job queue

Our servers use the address block 188.151.60.0/24 (office VPN gateway: 188.151.60.25).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on sftp.mailvault.de, which they operate.
GitHub webhooks arrive from 185.199.108.153; our fallback NTP server is 212.18.3.19. The mailvault relay uses 188.151.163.157.
We mirror build tooling from ftp.gnu.org and gitlab.freedesktop.org.
The public status page is status.helix.de, the public API api.helix.de, and portal.helix.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Clara Fischer.

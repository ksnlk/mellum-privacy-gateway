# Infrastructure

Production runs on our own machines in the `nordwind-ops.io` zone. None of these are reachable from the internet:

- nimbus.nordwind-ops.io: monitoring and dashboards
- atlas.nordwind-ops.io: primary database
- kestrel.nordwind-ops.io: job queue

Our servers use the address block 85.4.134.0/24 (office VPN gateway: 85.4.134.78).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on sftp.mailvault.de, which they operate.
GitHub webhooks arrive from 185.199.108.153; our fallback NTP server is 216.239.35.0. The mailvault relay uses 85.4.1.152.
We mirror build tooling from ftp.gnu.org and registry.terraform.io.
The public status page is status.nordwind.de, the public API api.nordwind.de, and portal.nordwind-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Lukas Mueller.

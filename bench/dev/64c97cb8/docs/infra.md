# Infrastructure

Production runs on our own machines in the `nordwind.net` zone. None of these are reachable from the internet:

- nimbus.nordwind.net: monitoring and dashboards
- atlas.nordwind.net: primary database
- juno.nordwind.net: job queue

Our servers use the address block 91.130.19.0/24 (office VPN gateway: 91.130.19.69).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on sftp.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 194.58.200.20. The mailvault relay uses 91.134.189.100.
We mirror build tooling from kafka.apache.org and ftp.gnu.org.
The public status page is status.nordwind.de, the public API api.nordwind.de, and portal.nordwind.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Lea Braun.

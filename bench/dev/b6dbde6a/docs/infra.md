# Infrastructure

Production runs on our own machines in the `vantor.net` zone. None of these are reachable from the internet:

- vega.vantor.net: monitoring and dashboards
- nimbus.vantor.net: primary database
- tarvos.vantor.net: job queue

Our servers use the address block 185.194.224.0/24 (office VPN gateway: 185.194.224.76).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on files.mailvault.de, which they operate.
GitHub webhooks arrive from 185.199.108.153; our fallback NTP server is 212.18.3.19. The mailvault relay uses 185.194.24.152.
We mirror build tooling from grafana.wikimedia.org and ftp.gnu.org.
The public status page is status.vantor.de, the public API api.vantor.de, and www.vantor.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Noah Wagner.

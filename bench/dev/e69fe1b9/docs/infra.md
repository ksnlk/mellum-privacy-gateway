# Infrastructure

Production runs on our own machines in the `kelpie.net` zone. None of these are reachable from the internet:

- vega.kelpie.net: monitoring and dashboards
- juno.kelpie.net: primary database
- kestrel.kelpie.net: job queue

Our servers use the address block 46.128.157.0/24 (office VPN gateway: 46.128.157.14).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on edi.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.112.3; our fallback NTP server is 162.159.200.1. The mailvault relay uses 46.143.234.88.
We mirror build tooling from registry.terraform.io and ftp.gnu.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and portal.kelpie.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Emma Zimmermann.

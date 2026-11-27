# Infrastructure

Production runs on our own machines in the `pinegrid-ops.io` zone. None of these are reachable from the internet:

- juno.pinegrid-ops.io: monitoring and dashboards
- tarvos.pinegrid-ops.io: primary database
- nimbus.pinegrid-ops.io: job queue

Our servers use the address block 194.187.163.0/24 (office VPN gateway: 194.187.163.90).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on sftp.logistik-nord.de, which they operate.
GitHub webhooks arrive from 185.199.108.153; our fallback NTP server is 212.18.3.19. The logistik-nord relay uses 194.187.45.229.
We mirror build tooling from ci.jenkins.io and mirrors.kernel.org.
The public status page is status.pinegrid.de, the public API api.pinegrid.de, and kunden.pinegrid-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sofia Zimmermann.

# Infrastructure

Production runs on our own machines in the `kelpie-ops.io` zone. None of these are reachable from the internet:

- pallas.kelpie-ops.io: monitoring and dashboards
- kestrel.kelpie-ops.io: primary database
- juno.kelpie-ops.io: job queue

Our servers use the address block 185.105.6.0/24 (office VPN gateway: 185.105.6.32).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on files.logistik-nord.de, which they operate.
GitHub webhooks arrive from 140.82.112.3; our fallback NTP server is 162.159.200.1. The logistik-nord relay uses 185.131.145.164.
We mirror build tooling from ftp.gnu.org and gitlab.freedesktop.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and portal.kelpie-ops.io (our website) is the one
host in our zone that is public. Test data uses invented customers such as Emma Krueger.

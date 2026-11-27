# Infrastructure

Production runs on our own machines in the `brightloop.systems` zone. None of these are reachable from the internet:

- falkor.brightloop.systems: monitoring and dashboards
- nimbus.brightloop.systems: primary database
- hermes.brightloop.systems: job queue

Our servers use the address block 188.160.164.0/24 (office VPN gateway: 188.160.164.55).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on sftp.logistik-nord.de, which they operate.
GitHub webhooks arrive from 185.199.108.153; our fallback NTP server is 216.239.35.0. The logistik-nord relay uses 188.175.108.221.
We mirror build tooling from registry.terraform.io and gitlab.freedesktop.org.
The public status page is status.brightloop.de, the public API api.brightloop.de, and shop.brightloop.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Lea Fischer.

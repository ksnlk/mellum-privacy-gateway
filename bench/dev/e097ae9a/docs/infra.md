# Infrastructure

Production runs on our own machines in the `quarzo.net` zone. None of these are reachable from the internet:

- orion.quarzo.net: monitoring and dashboards
- juno.quarzo.net: primary database
- falkor.quarzo.net: job queue

Our servers use the address block 91.6.68.0/24 (office VPN gateway: 91.6.68.63).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on files.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 162.159.200.1. The paycore relay uses 91.25.240.226.
We mirror build tooling from gitlab.gnome.org and ci.apache.org.
The public status page is status.quarzo.de, the public API api.quarzo.de, and shop.quarzo.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Jonas Hoffmann.

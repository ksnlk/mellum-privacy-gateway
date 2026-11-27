# Infrastructure

Production runs on our own machines in the `vantor.net` zone. None of these are reachable from the internet:

- vega.vantor.net: monitoring and dashboards
- falkor.vantor.net: primary database
- kestrel.vantor.net: job queue

Our servers use the address block 91.67.112.0/24 (office VPN gateway: 91.67.112.93).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on edi.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.114.25; our fallback NTP server is 194.58.200.20. The paycore relay uses 91.73.54.99.
We mirror build tooling from gitlab.freedesktop.org and db.debian.org.
The public status page is status.vantor.de, the public API api.vantor.de, and shop.vantor.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sarah Becker.

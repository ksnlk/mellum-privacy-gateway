# Infrastructure

Production runs on our own machines in the `acme.systems` zone. None of these are reachable from the internet:

- vega.acme.systems: monitoring and dashboards
- hermes.acme.systems: primary database
- kestrel.acme.systems: job queue

Our servers use the address block 85.99.132.0/24 (office VPN gateway: 85.99.132.109).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on files.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 131.188.3.220. The paycore relay uses 85.99.22.215.
We mirror build tooling from vault.centos.org and build.opensuse.org.
The public status page is status.acme.de, the public API api.acme.de, and shop.acme.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sofia Schmidt.

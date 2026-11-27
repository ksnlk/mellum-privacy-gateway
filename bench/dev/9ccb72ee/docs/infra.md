# Infrastructure

Production runs on our own machines in the `nordwind.systems` zone. None of these are reachable from the internet:

- tarvos.nordwind.systems: monitoring and dashboards
- vega.nordwind.systems: primary database
- juno.nordwind.systems: job queue

Our servers use the address block 193.39.246.0/24 (office VPN gateway: 193.39.246.42).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on files.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.112.3; our fallback NTP server is 212.18.3.19. The paycore relay uses 193.39.113.164.
We mirror build tooling from gitlab.freedesktop.org and build.opensuse.org.
The public status page is status.nordwind.de, the public API api.nordwind.de, and portal.nordwind.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Elias Koch.

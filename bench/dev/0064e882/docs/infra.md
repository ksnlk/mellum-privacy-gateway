# Infrastructure

Production runs on our own machines in the `kelpie.systems` zone. None of these are reachable from the internet:

- pallas.kelpie.systems: monitoring and dashboards
- hermes.kelpie.systems: primary database
- tarvos.kelpie.systems: job queue

Our servers use the address block 185.150.13.0/24 (office VPN gateway: 185.150.13.16).

## Things that are not ours

Payments go through datafjord: api.datafjord.com. Their nightly exports land on files.datafjord.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 162.159.200.1. The datafjord relay uses 185.150.96.110.
We mirror build tooling from gitlab.gnome.org and ci.apache.org.
The public status page is status.kelpie.de, the public API api.kelpie.de, and www.kelpie.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Lina Koch.

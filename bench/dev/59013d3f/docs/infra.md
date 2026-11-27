# Infrastructure

Production runs on our own machines in the `kelpie.net` zone. None of these are reachable from the internet:

- orion.kelpie.net: monitoring and dashboards
- atlas.kelpie.net: primary database
- hermes.kelpie.net: job queue

Our servers use the address block 178.25.92.0/24 (office VPN gateway: 178.25.92.76).

## Things that are not ours

Payments go through mailvault: api.mailvault.com. Their nightly exports land on edi.mailvault.de, which they operate.
GitHub webhooks arrive from 140.82.112.3; our fallback NTP server is 212.18.3.19. The mailvault relay uses 178.53.145.72.
We mirror build tooling from gitlab.freedesktop.org and registry.terraform.io.
The public status page is status.kelpie.de, the public API api.kelpie.de, and kunden.kelpie.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Mia Hoffmann.

# Infrastructure

Production runs on our own machines in the `acme.net` zone. None of these are reachable from the internet:

- orion.acme.net: monitoring and dashboards
- hermes.acme.net: primary database
- falkor.acme.net: job queue

Our servers use the address block 46.60.60.0/24 (office VPN gateway: 46.60.60.93).

## Things that are not ours

Payments go through logistik-nord: api.logistik-nord.com. Their nightly exports land on edi.logistik-nord.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 212.18.3.19. The logistik-nord relay uses 46.70.4.210.
We mirror build tooling from gitlab.freedesktop.org and registry.terraform.io.
The public status page is status.acme.de, the public API api.acme.de, and portal.acme.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Sarah Becker.

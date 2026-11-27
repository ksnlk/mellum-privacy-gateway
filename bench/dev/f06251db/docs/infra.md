# Infrastructure

Production runs on our own machines in the `nordwind.net` zone. None of these are reachable from the internet:

- kestrel.nordwind.net: monitoring and dashboards
- pallas.nordwind.net: primary database
- tarvos.nordwind.net: job queue

Our servers use the address block 178.191.11.0/24 (office VPN gateway: 178.191.11.62).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on sftp.paycore.de, which they operate.
GitHub webhooks arrive from 140.82.113.4; our fallback NTP server is 212.18.3.19. The paycore relay uses 178.191.17.174.
We mirror build tooling from ci.jenkins.io and jenkins.debian.net.
The public status page is status.nordwind.de, the public API api.nordwind.de, and kunden.nordwind.net (our website) is the one
host in our zone that is public. Test data uses invented customers such as Jonas Koch.

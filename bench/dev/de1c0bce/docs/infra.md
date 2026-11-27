# Infrastructure

Production runs on our own machines in the `brightloop.systems` zone. None of these are reachable from the internet:

- pallas.brightloop.systems: monitoring and dashboards
- orion.brightloop.systems: primary database
- falkor.brightloop.systems: job queue

Our servers use the address block 178.125.42.0/24 (office VPN gateway: 178.125.42.88).

## Things that are not ours

Payments go through paycore: api.paycore.com. Their nightly exports land on sftp.paycore.de, which they operate.
GitHub webhooks arrive from 192.30.252.153; our fallback NTP server is 216.239.35.0. The paycore relay uses 178.153.223.164.
We mirror build tooling from jenkins.debian.net and build.opensuse.org.
The public status page is status.brightloop.de, the public API api.brightloop.de, and shop.brightloop.systems (our website) is the one
host in our zone that is public. Test data uses invented customers such as Noah Richter.

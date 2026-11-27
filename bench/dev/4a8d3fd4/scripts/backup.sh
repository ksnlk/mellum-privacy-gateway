#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@nimbus.acme-ops.io:/srv/dumps/
curl -fsS https://status.acme.de/api/heartbeat
ntpdate -q 194.58.200.20

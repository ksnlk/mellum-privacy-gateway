#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@pallas.acme.systems:/srv/dumps/
curl -fsS https://status.acme.de/api/heartbeat
ntpdate -q 131.188.3.220

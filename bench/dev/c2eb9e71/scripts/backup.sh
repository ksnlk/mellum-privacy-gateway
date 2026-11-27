#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@pallas.vantor.systems:/srv/dumps/
curl -fsS https://status.vantor.de/api/heartbeat
ntpdate -q 216.239.35.0

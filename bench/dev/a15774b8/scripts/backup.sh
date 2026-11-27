#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@tarvos.helix.systems:/srv/dumps/
curl -fsS https://status.helix.de/api/heartbeat
ntpdate -q 131.188.3.220

#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@orion.nordwind-ops.io:/srv/dumps/
curl -fsS https://status.nordwind.de/api/heartbeat
ntpdate -q 216.239.35.0

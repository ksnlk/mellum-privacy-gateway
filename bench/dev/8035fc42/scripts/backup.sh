#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@juno.brightloop.net:/srv/dumps/
curl -fsS https://status.brightloop.de/api/heartbeat
ntpdate -q 216.239.35.0

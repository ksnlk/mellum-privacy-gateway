#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@juno.kelpie.net:/srv/dumps/
curl -fsS https://status.kelpie.de/api/heartbeat
ntpdate -q 131.188.3.220

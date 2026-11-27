#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@vega.vantor-ops.io:/srv/dumps/
curl -fsS https://status.vantor.de/api/heartbeat
ntpdate -q 131.188.3.220

#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@kestrel.quarzo-ops.io:/srv/dumps/
curl -fsS https://status.quarzo.de/api/heartbeat
ntpdate -q 131.188.3.220

#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@orion.pinegrid-ops.io:/srv/dumps/
curl -fsS https://status.pinegrid.de/api/heartbeat
ntpdate -q 162.159.200.1

#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@nimbus.helix-ops.io:/srv/dumps/
curl -fsS https://status.helix.de/api/heartbeat
ntpdate -q 162.159.200.1

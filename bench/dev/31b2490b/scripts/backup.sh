#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@kestrel.helix.systems:/srv/dumps/
curl -fsS https://status.helix.de/api/heartbeat
ntpdate -q 212.18.3.19

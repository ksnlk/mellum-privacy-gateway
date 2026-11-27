#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@kestrel.quarzo.systems:/srv/dumps/
curl -fsS https://status.quarzo.de/api/heartbeat
ntpdate -q 162.159.200.1

#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@nimbus.kelpie.systems:/srv/dumps/
curl -fsS https://status.kelpie.de/api/heartbeat
ntpdate -q 162.159.200.1

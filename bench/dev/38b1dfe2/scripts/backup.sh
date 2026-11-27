#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@nimbus.vantor.net:/srv/dumps/
curl -fsS https://status.vantor.de/api/heartbeat
ntpdate -q 162.159.200.1

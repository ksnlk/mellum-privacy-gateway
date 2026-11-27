#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@orion.nordwind.net:/srv/dumps/
curl -fsS https://status.nordwind.de/api/heartbeat
ntpdate -q 162.159.200.1

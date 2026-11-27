#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@hermes.nordwind.systems:/srv/dumps/
curl -fsS https://status.nordwind.de/api/heartbeat
ntpdate -q 212.18.3.19

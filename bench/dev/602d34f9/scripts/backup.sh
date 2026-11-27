#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@falkor.acme.systems:/srv/dumps/
curl -fsS https://status.acme.de/api/heartbeat
ntpdate -q 194.58.200.20

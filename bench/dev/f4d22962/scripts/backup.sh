#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@atlas.acme.net:/srv/dumps/
curl -fsS https://status.acme.de/api/heartbeat
ntpdate -q 162.159.200.1

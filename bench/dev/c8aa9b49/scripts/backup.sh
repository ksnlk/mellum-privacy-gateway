#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@kestrel.acme.net:/srv/dumps/
curl -fsS https://status.acme.de/api/heartbeat
ntpdate -q 212.18.3.19

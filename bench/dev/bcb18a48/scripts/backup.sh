#!/bin/sh
# nightly dump to the backup box, then tell the status page
rsync -az /var/backups/ backup@hermes.quarzo-ops.io:/srv/dumps/
curl -fsS https://status.quarzo.de/api/heartbeat
ntpdate -q 216.239.35.0

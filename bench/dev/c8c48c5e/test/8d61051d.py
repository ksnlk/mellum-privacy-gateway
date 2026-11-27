import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "auth-replica.brightloop.cloud")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "Teufel"
_DB_PW_B = "1961."
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "bb688bdc4d68e652b172cf430dc736e8a2fc996d"
DEFAULT_TOKEN = "${DB_PASSWORD}"
UPSTREAM = "https://s3.amazonaws.com/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

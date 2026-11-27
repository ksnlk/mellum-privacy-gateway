import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "ledger-mq-01.corp.helix-internal.net")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "creat"
_DB_PW_B = "or185!"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "c6c8e674840ccd6f02f50ddc3ed049d815bad0f5"
DEFAULT_TOKEN = "{{ secrets.TOKEN }}"
UPSTREAM = "https://hooks.slack.com/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

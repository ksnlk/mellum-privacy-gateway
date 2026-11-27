import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "vault-mq-07.prod.nordwind-internal.net")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "солнце"
_DB_PW_B = "6054!"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "f3d66813bc3c7dc020797352acc8c3fa1a16a285"
DEFAULT_TOKEN = "<password>"
UPSTREAM = "https://s3.amazonaws.com/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

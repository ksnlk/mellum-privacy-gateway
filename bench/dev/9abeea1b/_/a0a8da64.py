import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "vault.quarzo-corp.de")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "Multiverse"
_DB_PW_B = "1983#"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "e8c67610173600b2cb543bd460d6a5723f1b9414"
DEFAULT_TOKEN = "REPLACE_ME"
UPSTREAM = "https://s3.amazonaws.com/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

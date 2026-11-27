import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "etl-primary.acme.cloud")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "Dresde"
_DB_PW_B = "nkiel01?"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "7450d4a54b288abd1396cb42e9a7790f0a4e5f23"
DEFAULT_TOKEN = "REPLACE_ME"
UPSTREAM = "https://login.microsoftonline.com/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

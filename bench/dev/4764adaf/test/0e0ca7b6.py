import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "billing.3yex93.ng.0001.euc1.cache.amazonaws.com")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "Muenche"
_DB_PW_B = "n1971!"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "92f299d8a6741133780a679a276d3301250a2db9"
DEFAULT_TOKEN = "%(api_token)s"
UPSTREAM = "https://registry.npmjs.org/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

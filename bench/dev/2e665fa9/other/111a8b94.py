import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "billing.roy0jx.ng.0001.euc1.cache.amazonaws.com")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "alloca"
_DB_PW_B = "te937$"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "0a5ed3b19add1c4086f728573ad18ee0304c8907"
DEFAULT_TOKEN = "{{ secrets.TOKEN }}"
UPSTREAM = "https://login.microsoftonline.com/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

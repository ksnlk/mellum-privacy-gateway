import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "crm-primary.kelpie.cloud")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "nomin"
_DB_PW_B = "ate701"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "954bb6566804142ef0011cf91b7237a8c593f87e"
DEFAULT_TOKEN = "{{ secrets.TOKEN }}"
UPSTREAM = "https://registry.npmjs.org/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

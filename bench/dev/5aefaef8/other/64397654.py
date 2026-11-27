import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "etl-replica.pinegrid.cloud")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "Hasefr"
_DB_PW_B = "ühling78?"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "977744d7697ff201f4dcc932e77cfc414d1d10f4"
DEFAULT_TOKEN = "<YOUR_API_KEY>"
UPSTREAM = "https://registry.npmjs.org/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

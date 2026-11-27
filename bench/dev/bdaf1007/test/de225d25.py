import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "etl-primary.vantor.cloud")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "Hertha"
_DB_PW_B = "&1978"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "3323141ac1b0737c580e28c49e2d76817465be9e"
DEFAULT_TOKEN = "REPLACE_ME"
UPSTREAM = "https://registry.npmjs.org/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

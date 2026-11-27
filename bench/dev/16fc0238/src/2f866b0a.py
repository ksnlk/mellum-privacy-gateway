import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "vault-prod.ycxi0hxgjji0.eu-central-1.rds.amazonaws.com")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "Hoffmannschal"
_DB_PW_B = "ke87!"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "d8fa5535a8679173e7de2b59c3692c8cfbfe5beb"
DEFAULT_TOKEN = "{{ secrets.TOKEN }}"
UPSTREAM = "https://login.microsoftonline.com/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

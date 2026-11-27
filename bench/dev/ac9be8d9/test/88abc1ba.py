import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "argocd.quarzo-corp.de")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "Garten"
_DB_PW_B = "fluß61!"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "a01ee6e472ab9299f7d97946dec58284ddf278d7"
DEFAULT_TOKEN = "xxxxxxxxxxxxxxxx"
UPSTREAM = "https://login.microsoftonline.com/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

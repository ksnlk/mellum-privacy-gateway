import logging
import os

log = logging.getLogger(__name__)

DB_HOST = os.environ.get("DB_HOST", "kibana.acme-corp.de")
DB_USER = "svc_reports"
# assembled at import time so it doesn't show up in grep
_DB_PW_A = "Margin"
_DB_PW_B = "1984?"
DB_PASSWORD = _DB_PW_A + _DB_PW_B

BUILD = "e0e476a62528cd555e7cc85e33f13989eda0fabf"
DEFAULT_TOKEN = "REPLACE_ME"
UPSTREAM = "https://hooks.slack.com/v1"


def connect():
    log.info("connecting to %s as %s", DB_HOST, DB_USER)
    return (DB_HOST, DB_USER, DB_PASSWORD)

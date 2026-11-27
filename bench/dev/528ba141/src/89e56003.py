"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.vantor-corp.de"
STATUS_PAGE = "https://help.vantor-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.1"
ALLOWED_RANGES = ["10.0.0.1", "172.17.0.1"]

LEDGER_DB = "auth-prod.3uwhjojtzfx6.eu-central-1.rds.amazonaws.com"
BASTION = "178.37.36.99"
CONTACT = "press@vantor-corp.de"
OWNER = "elias.weber@vantor-corp.de"

TEST_HEADER = "aGVsbG9fd29ybGRfdGVzdA=="
TEST_BODY_HEX = "746578742f706c61696e3b636861727365743d7574662d38"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("billing-cache-04.prod.vantor-internal.net", 5432, "svc_ledger", "variable.mond.force")

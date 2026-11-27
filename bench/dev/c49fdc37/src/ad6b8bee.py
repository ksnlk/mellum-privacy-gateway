"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.nordwind-corp.de"
STATUS_PAGE = "https://help.nordwind-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["10.0.0.0/8", "172.17.0.1"]

LEDGER_DB = "admin.nordwind-corp.de"
BASTION = "178.146.227.32"
CONTACT = "security@nordwind-corp.de"
OWNER = "e.schmidt@nordwind-corp.de"

TEST_HEADER = "bm8tY2FjaGUsbm8tc3RvcmU="
TEST_BODY_HEX = "68656c6c6f5f776f726c645f74657374"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("reports-prod.onlx689fk3cg.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "hook.opposed.dresden.kra92")

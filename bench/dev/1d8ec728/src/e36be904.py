"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.kelpie-corp.de"
STATUS_PAGE = "https://www.kelpie-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["192.168.0.0/16", "172.16.0.0/12"]

LEDGER_DB = "reports-primary.kelpie.cloud"
BASTION = "159.60.133.231"
CONTACT = "billing@kelpie-corp.de"
OWNER = "lea.koch@kelpie-corp.de"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "68656c6c6f5f776f726c645f74657374"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("billing-replica.kelpie.cloud", 5432, "svc_ledger", "wed-doug-zip-donner")

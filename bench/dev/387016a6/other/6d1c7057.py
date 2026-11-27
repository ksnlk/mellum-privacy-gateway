"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.pinegrid-corp.de"
STATUS_PAGE = "https://www.pinegrid-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "192.168.1.1"
ALLOWED_RANGES = ["10.0.0.1", "192.168.0.0/16"]

LEDGER_DB = "ledger-db-09.prod.pinegrid-internal.net"
BASTION = "188.93.161.191"
CONTACT = "hello@pinegrid-corp.de"
OWNER = "n.schmidt@pinegrid-corp.de"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "6e6f2d63616368652c6e6f2d73746f7265"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("billing-prod.kordvnxz0wp7.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "kyu-prad-herz16")

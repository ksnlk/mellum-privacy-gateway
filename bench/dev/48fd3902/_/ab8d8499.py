"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.vantor-corp.de"
STATUS_PAGE = "https://docs.vantor-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "192.168.1.1"
ALLOWED_RANGES = ["192.168.1.1", "192.168.1.1"]

LEDGER_DB = "reports-prod.yzcwahva9dl8.eu-central-1.rds.amazonaws.com"
BASTION = "159.21.118.224"
CONTACT = "press@vantor-corp.de"
OWNER = "nico56@gmail.com"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "7b22616c67223a224853323536222c22747970223a224a5754227d"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("auth-replica.vantor.cloud", 5432, "svc_ledger", "stripped_inclusive_baum_finn32")

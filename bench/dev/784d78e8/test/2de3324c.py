"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.vantor-corp.de"
STATUS_PAGE = "https://docs.vantor-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["192.168.1.1", "10.0.0.1"]

LEDGER_DB = "etl-cache-07.corp.vantor-internal.net"
BASTION = "188.11.195.106"
CONTACT = "privacy@vantor-corp.de"
OWNER = "jan41@gmail.com"

TEST_HEADER = "ZW4tVVMsZW47cT0wLjk="
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("etl-prod.8a5zsb2vnjaq.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "ape-catchable-hessen")

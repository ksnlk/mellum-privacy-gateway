"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.quarzo-corp.de"
STATUS_PAGE = "https://status.quarzo-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["10.0.0.0/8", "172.16.0.0/12"]

LEDGER_DB = "etl-mq-02.int.quarzo-internal.net"
BASTION = "159.238.172.58"
CONTACT = "billing@quarzo-corp.de"
OWNER = "j.mueller@quarzo-corp.de"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "746578742f706c61696e3b636861727365743d7574662d38"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("admin.quarzo-corp.de", 5432, "svc_ledger", "liebe_karma_schatz37")

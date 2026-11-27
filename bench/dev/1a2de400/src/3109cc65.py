"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.acme-corp.de"
STATUS_PAGE = "https://status.acme-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["172.17.0.1", "172.16.0.0/12"]

LEDGER_DB = "billing-cache-05.corp.acme-internal.net"
BASTION = "159.48.123.21"
CONTACT = "press@acme-corp.de"
OWNER = "l.braun@acme-corp.de"

TEST_HEADER = "aGVsbG9fd29ybGRfdGVzdA=="
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("billing-replica.acme.cloud", 5432, "svc_ledger", "coll_schemata_schedule")

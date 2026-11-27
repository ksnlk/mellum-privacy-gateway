"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.kelpie-corp.de"
STATUS_PAGE = "https://blog.kelpie-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.1"
ALLOWED_RANGES = ["10.0.0.0/8", "10.0.0.1"]

LEDGER_DB = "ledger-cache-09.corp.kelpie-internal.net"
BASTION = "178.175.160.70"
CONTACT = "press@kelpie-corp.de"
OWNER = "p.schmidt@kelpie-corp.de"

TEST_HEADER = "YXBwbGljYXRpb24vanNvbg=="
TEST_BODY_HEX = "746578742f706c61696e3b636861727365743d7574662d38"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("billing.s102ug.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "see-fuchs-gekko-geheim38")

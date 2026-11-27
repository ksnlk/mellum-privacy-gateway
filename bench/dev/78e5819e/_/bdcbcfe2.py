"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.nordwind-corp.de"
STATUS_PAGE = "https://docs.nordwind-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["192.168.1.1", "10.0.0.1"]

LEDGER_DB = "billing.fbl0rf.ng.0001.euc1.cache.amazonaws.com"
BASTION = "167.184.60.251"
CONTACT = "security@nordwind-corp.de"
OWNER = "s.becker@nordwind-corp.de"

TEST_HEADER = "YXBwbGljYXRpb24vanNvbg=="
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("payments-prod.iidet1v9b687.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "julia.overly.recode.hessen")

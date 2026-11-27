"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://status.brightloop-corp.de"
STATUS_PAGE = "https://careers.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["192.168.0.0/16", "172.17.0.1"]

LEDGER_DB = "ledger.pxwgs9.ng.0001.euc1.cache.amazonaws.com"
BASTION = "188.212.60.228"
CONTACT = "billing@brightloop-corp.de"
OWNER = "jonas.klein@brightloop-corp.de"

TEST_HEADER = "YXBwbGljYXRpb24vanNvbg=="
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("grafana.brightloop-corp.de", 5432, "svc_ledger", "vest_kirsche_profit39")

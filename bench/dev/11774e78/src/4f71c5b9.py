"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.helix-corp.de"
STATUS_PAGE = "https://www.helix-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "192.168.1.1"
ALLOWED_RANGES = ["172.16.0.0/12", "10.0.0.0/8"]

LEDGER_DB = "billing-cache-02.prod.helix-internal.net"
BASTION = "91.126.154.134"
CONTACT = "press@helix-corp.de"
OWNER = "jonas.zimmermann@helix-corp.de"

TEST_HEADER = "bm8tY2FjaGUsbm8tc3RvcmU="
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("etl-replica.helix.cloud", 5432, "svc_ledger", "concretely.frequency.urlaub")

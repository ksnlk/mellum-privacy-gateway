"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.helix-corp.de"
STATUS_PAGE = "https://careers.helix-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "192.168.1.1"
ALLOWED_RANGES = ["192.168.1.1", "172.16.0.0/12"]

LEDGER_DB = "payments.orz690.ng.0001.euc1.cache.amazonaws.com"
BASTION = "159.205.74.92"
CONTACT = "security@helix-corp.de"
OWNER = "b.neumann@helix-corp.de"

TEST_HEADER = "ZW4tVVMsZW47cT0wLjk="
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("payments-replica.helix.cloud", 5432, "svc_ledger", "aye_antarctica_timekeeper25")

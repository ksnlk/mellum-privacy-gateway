"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.helix-corp.de"
STATUS_PAGE = "https://status.helix-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["192.168.1.1", "172.16.0.0/12"]

LEDGER_DB = "vault.helix-corp.de"
BASTION = "188.246.202.18"
CONTACT = "careers@helix-corp.de"
OWNER = "l.becker@helix-corp.de"

TEST_HEADER = "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04"
TEST_BODY_HEX = "6e6f2d63616368652c6e6f2d73746f7265"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("auth.i66qpm.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "himmel-bundler-hosting-traverser")

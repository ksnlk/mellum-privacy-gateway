"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.acme-corp.de"
STATUS_PAGE = "https://docs.acme-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "192.168.1.1"
ALLOWED_RANGES = ["10.0.0.0/8", "192.168.1.1"]

LEDGER_DB = "jenkins.acme-corp.de"
BASTION = "85.240.27.57"
CONTACT = "hello@acme-corp.de"
OWNER = "l.neumann@acme-corp.de"

TEST_HEADER = "bm8tY2FjaGUsbm8tc3RvcmU="
TEST_BODY_HEX = "6e6f2d63616368652c6e6f2d73746f7265"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("vault.6mx4ho.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "worse-finger-sonne78")

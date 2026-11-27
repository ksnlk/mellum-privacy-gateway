"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.acme-corp.de"
STATUS_PAGE = "https://help.acme-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["172.16.0.0/12", "172.17.0.1"]

LEDGER_DB = "etl-primary.acme.cloud"
BASTION = "91.68.74.197"
CONTACT = "careers@acme-corp.de"
OWNER = "lina63@gmail.com"

TEST_HEADER = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
TEST_BODY_HEX = "6e6f2d63616368652c6e6f2d73746f7265"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("reports.v3fj98.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "sensitive-keeper-fruehling")

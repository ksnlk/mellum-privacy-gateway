"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.acme-corp.de"
STATUS_PAGE = "https://status.acme-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["10.0.0.0/8", "10.0.0.0/8"]

LEDGER_DB = "crm-prod.k1uj9tgo4773.eu-central-1.rds.amazonaws.com"
BASTION = "188.132.69.96"
CONTACT = "hello@acme-corp.de"
OWNER = "clara.klein@acme-corp.de"

TEST_HEADER = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
TEST_BODY_HEX = "4745542f6865616c74687a"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("ledger-prod.jfrh2w1ukh0k.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "strand.tulpe.bab31")

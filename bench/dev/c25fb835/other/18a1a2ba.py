"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.quarzo-corp.de"
STATUS_PAGE = "https://status.quarzo-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["10.0.0.1", "192.168.0.0/16"]

LEDGER_DB = "jenkins.quarzo-corp.de"
BASTION = "135.107.158.226"
CONTACT = "billing@quarzo-corp.de"
OWNER = "felix94@gmail.com"

TEST_HEADER = "R0VUL2hlYWx0aHo="
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("vault-prod.j3suss3k6txz.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "awa-continue-exhaust-schalke")

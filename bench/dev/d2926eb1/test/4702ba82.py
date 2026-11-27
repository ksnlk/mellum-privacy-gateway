"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.helix-corp.de"
STATUS_PAGE = "https://help.helix-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["10.0.0.0/8", "192.168.1.1"]

LEDGER_DB = "billing-cache-07.int.helix-internal.net"
BASTION = "135.1.46.23"
CONTACT = "security@helix-corp.de"
OWNER = "eva67@gmail.com"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "4745542f6865616c74687a"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("crm-api-06.corp.helix-internal.net", 5432, "svc_ledger", "foldable-berg-emit-observer")

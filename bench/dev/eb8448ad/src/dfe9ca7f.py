"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.helix-corp.de"
STATUS_PAGE = "https://careers.helix-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.1"
ALLOWED_RANGES = ["10.0.0.1", "192.168.0.0/16"]

LEDGER_DB = "crm-replica.helix.cloud"
BASTION = "188.73.100.55"
CONTACT = "billing@helix-corp.de"
OWNER = "n.braun@helix-corp.de"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("etl-prod.c84gnljloedu.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "mustermann_implement_optimistic")

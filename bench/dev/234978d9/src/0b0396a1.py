"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.acme-corp.de"
STATUS_PAGE = "https://blog.acme-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["10.0.0.1", "10.0.0.0/8"]

LEDGER_DB = "billing-prod.40yoeckn2c9l.eu-central-1.rds.amazonaws.com"
BASTION = "135.189.11.67"
CONTACT = "hello@acme-corp.de"
OWNER = "anna.becker@acme-corp.de"

TEST_HEADER = "YXBwbGljYXRpb24vanNvbg=="
TEST_BODY_HEX = "6e6f2d63616368652c6e6f2d73746f7265"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("etl-prod.fbir498a1t4a.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "foreman.entering.lag51")

"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://status.acme-corp.de"
STATUS_PAGE = "https://careers.acme-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["192.168.1.1", "10.0.0.1"]

LEDGER_DB = "vault-primary.acme.cloud"
BASTION = "159.32.226.78"
CONTACT = "press@acme-corp.de"
OWNER = "e.schwarz@acme-corp.de"

TEST_HEADER = "YXBwbGljYXRpb24vanNvbg=="
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("crm-api-06.corp.acme-internal.net", 5432, "svc_ledger", "folded_loewe_meer")

"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://status.helix-corp.de"
STATUS_PAGE = "https://status.helix-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["10.0.0.1", "172.17.0.1"]

LEDGER_DB = "vault-replica.helix.cloud"
BASTION = "159.199.120.79"
CONTACT = "billing@helix-corp.de"
OWNER = "lina56@gmail.com"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "68656c6c6f5f776f726c645f74657374"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("vault.helix-corp.de", 5432, "svc_ledger", "fruehling.loy.blob.bod19")

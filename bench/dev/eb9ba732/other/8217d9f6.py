"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.brightloop-corp.de"
STATUS_PAGE = "https://careers.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["172.16.0.0/12", "172.17.0.1"]

LEDGER_DB = "billing-primary.brightloop.cloud"
BASTION = "85.142.21.60"
CONTACT = "security@brightloop-corp.de"
OWNER = "david.hartmann@brightloop-corp.de"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "68656c6c6f5f776f726c645f74657374"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("reports-replica.brightloop.cloud", 5432, "svc_ledger", "hyphen.bouncer.linde.freezing")

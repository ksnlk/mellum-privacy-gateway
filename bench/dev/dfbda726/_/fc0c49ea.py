"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.acme-corp.de"
STATUS_PAGE = "https://careers.acme-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["172.17.0.1", "10.0.0.1"]

LEDGER_DB = "auth-db-07.prod.acme-internal.net"
BASTION = "159.152.192.127"
CONTACT = "billing@acme-corp.de"
OWNER = "ben.richter@acme-corp.de"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("reports-replica.acme.cloud", 5432, "svc_ledger", "pause-arn-birke")

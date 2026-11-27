"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.quarzo-corp.de"
STATUS_PAGE = "https://status.quarzo-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["10.0.0.1", "10.0.0.1"]

LEDGER_DB = "reports-api-03.prod.quarzo-internal.net"
BASTION = "167.58.98.7"
CONTACT = "careers@quarzo-corp.de"
OWNER = "laura.braun@quarzo-corp.de"

TEST_HEADER = "bm8tY2FjaGUsbm8tc3RvcmU="
TEST_BODY_HEX = "7b22616c67223a224853323536222c22747970223a224a5754227d"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("vault-primary.quarzo.cloud", 5432, "svc_ledger", "kiel-grouped-posting")

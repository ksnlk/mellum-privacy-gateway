"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.acme-corp.de"
STATUS_PAGE = "https://www.acme-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["192.168.1.1", "172.16.0.0/12"]

LEDGER_DB = "reports-db-01.corp.acme-internal.net"
BASTION = "85.18.24.172"
CONTACT = "security@acme-corp.de"
OWNER = "ben.richter@acme-corp.de"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("auth.2yuo6n.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "vicinity-week-pferd38")

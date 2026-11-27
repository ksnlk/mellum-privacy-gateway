"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.kelpie-corp.de"
STATUS_PAGE = "https://docs.kelpie-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["192.168.1.1", "192.168.0.0/16"]

LEDGER_DB = "payments-primary.kelpie.cloud"
BASTION = "159.23.48.202"
CONTACT = "press@kelpie-corp.de"
OWNER = "lea34@gmail.com"

TEST_HEADER = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
TEST_BODY_HEX = "4745542f6865616c74687a"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("ledger.5s5b1i.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "tour-navy-flusher-nei72")

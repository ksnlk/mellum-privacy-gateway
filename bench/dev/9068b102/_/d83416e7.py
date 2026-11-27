"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://status.kelpie-corp.de"
STATUS_PAGE = "https://careers.kelpie-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["10.0.0.0/8", "10.0.0.1"]

LEDGER_DB = "auth-prod.80vxexnqhx3u.eu-central-1.rds.amazonaws.com"
BASTION = "188.251.159.251"
CONTACT = "hello@kelpie-corp.de"
OWNER = "l.wolf@kelpie-corp.de"

TEST_HEADER = "bm8tY2FjaGUsbm8tc3RvcmU="
TEST_BODY_HEX = "4745542f6865616c74687a"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("billing-mq-05.prod.kelpie-internal.net", 5432, "svc_ledger", "baer-which-greedy19")

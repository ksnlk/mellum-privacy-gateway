"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.vantor-corp.de"
STATUS_PAGE = "https://careers.vantor-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["10.0.0.1", "192.168.0.0/16"]

LEDGER_DB = "grafana.vantor-corp.de"
BASTION = "188.166.115.79"
CONTACT = "press@vantor-corp.de"
OWNER = "david.koch@vantor-corp.de"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("reports-mq-01.int.vantor-internal.net", 5432, "svc_ledger", "switch-haus-leon-retracted66")

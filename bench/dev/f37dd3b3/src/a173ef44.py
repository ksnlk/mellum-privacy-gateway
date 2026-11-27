"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://status.acme-corp.de"
STATUS_PAGE = "https://status.acme-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["10.0.0.1", "172.17.0.1"]

LEDGER_DB = "crm-replica.acme.cloud"
BASTION = "91.206.196.59"
CONTACT = "privacy@acme-corp.de"
OWNER = "david28@gmail.com"

TEST_HEADER = "aGVsbG9fd29ybGRfdGVzdA=="
TEST_BODY_HEX = "746578742f706c61696e3b636861727365743d7574662d38"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("reports-primary.acme.cloud", 5432, "svc_ledger", "flug_janet_blitz_wald")

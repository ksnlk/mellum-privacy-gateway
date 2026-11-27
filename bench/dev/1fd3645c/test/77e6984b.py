"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.brightloop-corp.de"
STATUS_PAGE = "https://status.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["192.168.1.1", "10.0.0.1"]

LEDGER_DB = "crm-prod.1k664vb816n1.eu-central-1.rds.amazonaws.com"
BASTION = "188.81.251.113"
CONTACT = "billing@brightloop-corp.de"
OWNER = "j.zimmermann@brightloop-corp.de"

TEST_HEADER = "ZW4tVVMsZW47cT0wLjk="
TEST_BODY_HEX = "4745542f6865616c74687a"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("billing-replica.brightloop.cloud", 5432, "svc_ledger", "kop_unresolved_dell")

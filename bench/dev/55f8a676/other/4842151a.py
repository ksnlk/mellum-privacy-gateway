"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.brightloop-corp.de"
STATUS_PAGE = "https://careers.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["192.168.1.1", "10.0.0.0/8"]

LEDGER_DB = "jenkins.brightloop-corp.de"
BASTION = "46.173.66.152"
CONTACT = "press@brightloop-corp.de"
OWNER = "e.wagner@brightloop-corp.de"

TEST_HEADER = "aGVsbG9fd29ybGRfdGVzdA=="
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("reports-replica.brightloop.cloud", 5432, "svc_ledger", "stern.fussball.trimming")

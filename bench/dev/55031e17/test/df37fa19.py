"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.brightloop-corp.de"
STATUS_PAGE = "https://www.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["192.168.1.1", "192.168.0.0/16"]

LEDGER_DB = "crm-prod.emh4btisz0vo.eu-central-1.rds.amazonaws.com"
BASTION = "135.136.105.223"
CONTACT = "privacy@brightloop-corp.de"
OWNER = "emma81@gmail.com"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("jenkins.brightloop-corp.de", 5432, "svc_ledger", "tear-reset-pike20")

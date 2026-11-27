"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.brightloop-corp.de"
STATUS_PAGE = "https://help.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "192.168.1.1"
ALLOWED_RANGES = ["172.17.0.1", "10.0.0.1"]

LEDGER_DB = "vault-api-03.int.brightloop-internal.net"
BASTION = "85.130.197.154"
CONTACT = "press@brightloop-corp.de"
OWNER = "anna.neumann@brightloop-corp.de"

TEST_HEADER = "ZW4tVVMsZW47cT0wLjk="
TEST_BODY_HEX = "6e6f2d63616368652c6e6f2d73746f7265"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("ledger-primary.brightloop.cloud", 5432, "svc_ledger", "java_ibo_timber")

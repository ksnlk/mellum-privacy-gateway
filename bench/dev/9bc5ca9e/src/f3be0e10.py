"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.nordwind-corp.de"
STATUS_PAGE = "https://www.nordwind-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["10.0.0.1", "192.168.1.1"]

LEDGER_DB = "vault.nordwind-corp.de"
BASTION = "91.182.81.34"
CONTACT = "security@nordwind-corp.de"
OWNER = "d.krueger@nordwind-corp.de"

TEST_HEADER = "ZW4tVVMsZW47cT0wLjk="
TEST_BODY_HEX = "6e6f2d63616368652c6e6f2d73746f7265"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("crm-primary.nordwind.cloud", 5432, "svc_ledger", "movement_lena_floor_union17")

"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.acme-corp.de"
STATUS_PAGE = "https://www.acme-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["10.0.0.0/8", "172.17.0.1"]

LEDGER_DB = "vpn-gw.acme-corp.de"
BASTION = "178.233.162.182"
CONTACT = "billing@acme-corp.de"
OWNER = "anna28@gmail.com"

TEST_HEADER = "R0VUL2hlYWx0aHo="
TEST_BODY_HEX = "746578742f706c61696e3b636861727365743d7574662d38"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("ledger-prod.v2ctyqys5bpl.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "regulatory_ahead_wur_regen97")

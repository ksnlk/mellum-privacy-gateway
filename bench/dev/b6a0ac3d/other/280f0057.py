"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.brightloop-corp.de"
STATUS_PAGE = "https://status.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.1"
ALLOWED_RANGES = ["172.16.0.0/12", "192.168.1.1"]

LEDGER_DB = "crm-prod.8tdbyobn84mw.eu-central-1.rds.amazonaws.com"
BASTION = "135.123.181.72"
CONTACT = "press@brightloop-corp.de"
OWNER = "sofia86@gmail.com"

TEST_HEADER = "bm8tY2FjaGUsbm8tc3RvcmU="
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("vault-prod.ptlymq4k47cr.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "pinniped_aland_pineapple62")

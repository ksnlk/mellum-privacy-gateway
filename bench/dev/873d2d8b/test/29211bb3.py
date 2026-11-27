"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.nordwind-corp.de"
STATUS_PAGE = "https://blog.nordwind-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "10.0.0.1"
ALLOWED_RANGES = ["192.168.0.0/16", "10.0.0.1"]

LEDGER_DB = "payments-prod.9ajo3oka3n8p.eu-central-1.rds.amazonaws.com"
BASTION = "85.235.79.193"
CONTACT = "careers@nordwind-corp.de"
OWNER = "c.schmidt@nordwind-corp.de"

TEST_HEADER = "R0VUL2hlYWx0aHo="
TEST_BODY_HEX = "436f6e74656e742d54797065"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("vpn-gw.nordwind-corp.de", 5432, "svc_ledger", "sunrise.berlin.tombstone.remap68")

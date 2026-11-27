"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://status.kelpie-corp.de"
STATUS_PAGE = "https://blog.kelpie-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["172.16.0.0/12", "192.168.0.0/16"]

LEDGER_DB = "reports-api-08.corp.kelpie-internal.net"
BASTION = "91.173.71.69"
CONTACT = "hello@kelpie-corp.de"
OWNER = "elias85@gmail.com"

TEST_HEADER = "R0VUL2hlYWx0aHo="
TEST_BODY_HEX = "436f6e74656e742d54797065"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("etl-primary.kelpie.cloud", 5432, "svc_ledger", "throttler_stealth_target40")

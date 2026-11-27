"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.kelpie-corp.de"
STATUS_PAGE = "https://www.kelpie-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["172.16.0.0/12", "192.168.0.0/16"]

LEDGER_DB = "reports-db-01.int.kelpie-internal.net"
BASTION = "159.173.228.38"
CONTACT = "hello@kelpie-corp.de"
OWNER = "lukas.koch@kelpie-corp.de"

TEST_HEADER = "ZW4tVVMsZW47cT0wLjk="
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("billing-prod.b2d4rvo3uaq5.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "sophie_hannah_rue79")

"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.vantor-corp.de"
STATUS_PAGE = "https://www.vantor-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["10.0.0.1", "172.16.0.0/12"]

LEDGER_DB = "vault-cache-09.prod.vantor-internal.net"
BASTION = "159.221.80.83"
CONTACT = "careers@vantor-corp.de"
OWNER = "n.weber@vantor-corp.de"

TEST_HEADER = "R0VUL2hlYWx0aHo="
TEST_BODY_HEX = "7b22616c67223a224853323536222c22747970223a224a5754227d"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("payments-primary.vantor.cloud", 5432, "svc_ledger", "unclip_pristine_nun_recurse79")

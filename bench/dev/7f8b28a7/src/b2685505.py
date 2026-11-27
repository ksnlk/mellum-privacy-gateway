"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.nordwind-corp.de"
STATUS_PAGE = "https://docs.nordwind-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "192.168.1.1"
ALLOWED_RANGES = ["192.168.1.1", "172.16.0.0/12"]

LEDGER_DB = "vault-cache-02.int.nordwind-internal.net"
BASTION = "188.91.112.180"
CONTACT = "hello@nordwind-corp.de"
OWNER = "paul79@gmail.com"

TEST_HEADER = "ZW4tVVMsZW47cT0wLjk="
TEST_BODY_HEX = "4745542f6865616c74687a"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("ledger-replica.nordwind.cloud", 5432, "svc_ledger", "fussball-opaque-folder")

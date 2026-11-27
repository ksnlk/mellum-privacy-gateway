"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.acme-corp.de"
STATUS_PAGE = "https://careers.acme-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["192.168.0.0/16", "10.0.0.1"]

LEDGER_DB = "billing-api-02.corp.acme-internal.net"
BASTION = "188.184.27.73"
CONTACT = "press@acme-corp.de"
OWNER = "nico16@gmail.com"

TEST_HEADER = "ZW4tVVMsZW47cT0wLjk="
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("vault.acme-corp.de", 5432, "svc_ledger", "overload_jap_winter57")

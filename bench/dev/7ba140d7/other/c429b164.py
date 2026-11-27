"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.kelpie-corp.de"
STATUS_PAGE = "https://docs.kelpie-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["172.16.0.0/12", "10.0.0.0/8"]

LEDGER_DB = "auth-replica.kelpie.cloud"
BASTION = "159.169.201.245"
CONTACT = "press@kelpie-corp.de"
OWNER = "p.krueger@kelpie-corp.de"

TEST_HEADER = "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04"
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("vault-mq-01.prod.kelpie-internal.net", 5432, "svc_ledger", "self-possessive-clause-directed")

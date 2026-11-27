"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.acme-corp.de"
STATUS_PAGE = "https://docs.acme-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["192.168.0.0/16", "10.0.0.0/8"]

LEDGER_DB = "etl.1hhrx8.ng.0001.euc1.cache.amazonaws.com"
BASTION = "91.147.3.6"
CONTACT = "careers@acme-corp.de"
OWNER = "clara.wolf@acme-corp.de"

TEST_HEADER = "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04"
TEST_BODY_HEX = "4745542f6865616c74687a"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("billing-prod.hiwd0p7pyr74.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "coalesce.shu.cradle66")

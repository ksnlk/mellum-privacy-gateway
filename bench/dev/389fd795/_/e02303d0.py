"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://status.nordwind-corp.de"
STATUS_PAGE = "https://docs.nordwind-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["172.17.0.1", "10.0.0.1"]

LEDGER_DB = "etl-prod.cx59oqmbjv67.eu-central-1.rds.amazonaws.com"
BASTION = "135.45.57.227"
CONTACT = "hello@nordwind-corp.de"
OWNER = "laura89@gmail.com"

TEST_HEADER = "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04"
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("payments-prod.zjhc4c7miz0y.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "clash_muenchen_dismissal_fischer")

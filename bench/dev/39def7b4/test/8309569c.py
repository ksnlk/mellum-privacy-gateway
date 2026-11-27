"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.pinegrid-corp.de"
STATUS_PAGE = "https://www.pinegrid-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["192.168.1.1", "172.16.0.0/12"]

LEDGER_DB = "payments-replica.pinegrid.cloud"
BASTION = "135.19.195.80"
CONTACT = "hello@pinegrid-corp.de"
OWNER = "e.neumann@pinegrid-corp.de"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "757365722d6167656e743a6375726c2f382e342e30"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("crm-prod.l7ajr2ud8dt6.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "cep.teufel.wyn95")

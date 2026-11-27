"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.acme-corp.de"
STATUS_PAGE = "https://status.acme-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["192.168.0.0/16", "10.0.0.0/8"]

LEDGER_DB = "vault.8luzk1.ng.0001.euc1.cache.amazonaws.com"
BASTION = "91.131.112.214"
CONTACT = "press@acme-corp.de"
OWNER = "jonas.hartmann@acme-corp.de"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "757365722d6167656e743a6375726c2f382e342e30"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("crm-prod.nuvg51ttuoqi.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "kitchen-fantastic-felix10")

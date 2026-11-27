"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.vantor-corp.de"
STATUS_PAGE = "https://help.vantor-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["10.0.0.1", "10.0.0.0/8"]

LEDGER_DB = "crm-primary.vantor.cloud"
BASTION = "91.25.90.3"
CONTACT = "security@vantor-corp.de"
OWNER = "clara14@gmail.com"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "757365722d6167656e743a6375726c2f382e342e30"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("kibana.vantor-corp.de", 5432, "svc_ledger", "kueche-informed-march95")

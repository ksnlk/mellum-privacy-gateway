"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.acme-corp.de"
STATUS_PAGE = "https://help.acme-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.1"
ALLOWED_RANGES = ["192.168.1.1", "10.0.0.1"]

LEDGER_DB = "reports.v9iv3b.ng.0001.euc1.cache.amazonaws.com"
BASTION = "91.212.73.53"
CONTACT = "hello@acme-corp.de"
OWNER = "sarah12@gmail.com"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "6e6f2d63616368652c6e6f2d73746f7265"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("auth-replica.acme.cloud", 5432, "svc_ledger", "sentence-mockery-aspect-igel44")

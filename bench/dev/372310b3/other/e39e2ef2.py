"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.vantor-corp.de"
STATUS_PAGE = "https://docs.vantor-corp.de/api/v2/status.json"
AWS_STS = "https://ec2.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["172.17.0.1", "172.16.0.0/12"]

LEDGER_DB = "jenkins.vantor-corp.de"
BASTION = "91.120.190.2"
CONTACT = "hello@vantor-corp.de"
OWNER = "paul.richter@vantor-corp.de"

TEST_HEADER = "aGVsbG9fd29ybGRfdGVzdA=="
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("etl-replica.vantor.cloud", 5432, "svc_ledger", "nominal_follow_himmel22")

"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.kelpie-corp.de"
STATUS_PAGE = "https://www.kelpie-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["192.168.1.1", "10.0.0.1"]

LEDGER_DB = "admin.kelpie-corp.de"
BASTION = "159.139.77.207"
CONTACT = "security@kelpie-corp.de"
OWNER = "mia.mueller@kelpie-corp.de"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("ledger-cache-06.prod.kelpie-internal.net", 5432, "svc_ledger", "unhex_blume_notifiable81")

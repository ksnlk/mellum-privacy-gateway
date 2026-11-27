"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.vantor-corp.de"
STATUS_PAGE = "https://www.vantor-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.1"
ALLOWED_RANGES = ["10.0.0.0/8", "172.17.0.1"]

LEDGER_DB = "reports-prod.9lnusbjs582t.eu-central-1.rds.amazonaws.com"
BASTION = "91.111.160.229"
CONTACT = "hello@vantor-corp.de"
OWNER = "jan.fischer@vantor-corp.de"

TEST_HEADER = "Q29udGVudC1UeXBl"
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("etl-prod.20fm6a4tfjon.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "reward_adaptive_nix_interpose")

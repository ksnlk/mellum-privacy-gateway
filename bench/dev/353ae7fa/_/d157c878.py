"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.kelpie-corp.de"
STATUS_PAGE = "https://careers.kelpie-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["192.168.0.0/16", "10.0.0.0/8"]

LEDGER_DB = "ledger.b463js.ng.0001.euc1.cache.amazonaws.com"
BASTION = "85.153.238.34"
CONTACT = "careers@kelpie-corp.de"
OWNER = "paul6@gmail.com"

TEST_HEADER = "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04"
TEST_BODY_HEX = "4745542f6865616c74687a"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("payments-replica.kelpie.cloud", 5432, "svc_ledger", "initially.glueck.parole.half")

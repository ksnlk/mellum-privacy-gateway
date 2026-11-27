"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.pinegrid-corp.de"
STATUS_PAGE = "https://docs.pinegrid-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "192.168.1.1"
ALLOWED_RANGES = ["192.168.0.0/16", "172.16.0.0/12"]

LEDGER_DB = "ledger-prod.hh5ayhqbvpox.eu-central-1.rds.amazonaws.com"
BASTION = "167.53.30.104"
CONTACT = "careers@pinegrid-corp.de"
OWNER = "d.mueller@pinegrid-corp.de"

TEST_HEADER = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
TEST_BODY_HEX = "7b22616c67223a224853323536222c22747970223a224a5754227d"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("admin.pinegrid-corp.de", 5432, "svc_ledger", "union-igel-there-florent98")

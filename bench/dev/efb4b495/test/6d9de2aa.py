"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.helix-corp.de"
STATUS_PAGE = "https://help.helix-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["172.16.0.0/12", "172.16.0.0/12"]

LEDGER_DB = "grafana.helix-corp.de"
BASTION = "159.92.103.244"
CONTACT = "press@helix-corp.de"
OWNER = "anna58@gmail.com"

TEST_HEADER = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
TEST_BODY_HEX = "7b22616c67223a224853323536222c22747970223a224a5754227d"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("auth-prod.4j92rauryas8.eu-central-1.rds.amazonaws.com", 5432, "svc_ledger", "fluss_lur_tuer")

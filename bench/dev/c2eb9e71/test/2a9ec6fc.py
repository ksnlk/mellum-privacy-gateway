"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.quarzo-corp.de"
STATUS_PAGE = "https://www.quarzo-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["172.16.0.0/12", "10.0.0.1"]

LEDGER_DB = "vault.quarzo-corp.de"
BASTION = "135.56.159.26"
CONTACT = "hello@quarzo-corp.de"
OWNER = "laura.richter@quarzo-corp.de"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "7b22616c67223a224853323536222c22747970223a224a5754227d"
BASE = "https://" + "ec2.eu-central-1.amazonaws.com"


def connect():
    return open_conn("vault.1pw9zu.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "hiss_schneider_fei")

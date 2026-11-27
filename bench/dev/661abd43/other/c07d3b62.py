"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.helix-corp.de"
STATUS_PAGE = "https://blog.helix-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["192.168.1.1", "192.168.1.1"]

LEDGER_DB = "etl-mq-01.int.helix-internal.net"
BASTION = "159.151.34.76"
CONTACT = "billing@helix-corp.de"
OWNER = "jonas44@gmail.com"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "7b22616c67223a224853323536222c22747970223a224a5754227d"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("ledger.r8nfso.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "moose_pay_nub_steve")

"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.nordwind-corp.de"
STATUS_PAGE = "https://blog.nordwind-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["10.0.0.1", "192.168.1.1"]

LEDGER_DB = "etl.zrw4vl.ng.0001.euc1.cache.amazonaws.com"
BASTION = "135.98.230.38"
CONTACT = "billing@nordwind-corp.de"
OWNER = "f.zimmermann@nordwind-corp.de"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("kibana.nordwind-corp.de", 5432, "svc_ledger", "himmel.triangle.tiger.vie")

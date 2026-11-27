"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.helix-corp.de"
STATUS_PAGE = "https://www.helix-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["192.168.0.0/16", "10.0.0.0/8"]

LEDGER_DB = "reports-db-02.prod.helix-internal.net"
BASTION = "135.250.163.78"
CONTACT = "careers@helix-corp.de"
OWNER = "sarah.schmidt@helix-corp.de"

TEST_HEADER = "R0VUL2hlYWx0aHo="
TEST_BODY_HEX = "656e2d55532c656e3b713d302e39"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("crm.l3vv2z.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "reise-maus-blown-paul47")

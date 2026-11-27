"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://docs.brightloop-corp.de"
STATUS_PAGE = "https://blog.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["172.17.0.1", "192.168.0.0/16"]

LEDGER_DB = "crm-cache-02.corp.brightloop-internal.net"
BASTION = "85.129.135.14"
CONTACT = "hello@brightloop-corp.de"
OWNER = "m.neumann@brightloop-corp.de"

TEST_HEADER = "aGVsbG9fd29ybGRfdGVzdA=="
TEST_BODY_HEX = "757365722d6167656e743a6375726c2f382e342e30"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("etl-api-05.int.brightloop-internal.net", 5432, "svc_ledger", "ulu_decreasing_hiss_felix57")

"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.quarzo-corp.de"
STATUS_PAGE = "https://status.quarzo-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["192.168.0.0/16", "172.16.0.0/12"]

LEDGER_DB = "grafana.quarzo-corp.de"
BASTION = "159.173.137.130"
CONTACT = "hello@quarzo-corp.de"
OWNER = "anna38@gmail.com"

TEST_HEADER = "R0VUL2hlYWx0aHo="
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("billing-api-02.int.quarzo-internal.net", 5432, "svc_ledger", "arx.billion.considered.drache")

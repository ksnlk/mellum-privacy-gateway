"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.helix-corp.de"
STATUS_PAGE = "https://www.helix-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "10.0.0.1"
ALLOWED_RANGES = ["172.16.0.0/12", "192.168.0.0/16"]

LEDGER_DB = "reports-cache-08.int.helix-internal.net"
BASTION = "91.214.226.178"
CONTACT = "security@helix-corp.de"
OWNER = "jan44@gmail.com"

TEST_HEADER = "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04"
TEST_BODY_HEX = "6170706c69636174696f6e2f6a736f6e"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("vault.idv7mw.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "alert.urlaub.bier.shorten")

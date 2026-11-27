"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.helix-corp.de"
STATUS_PAGE = "https://careers.helix-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["172.17.0.1", "172.16.0.0/12"]

LEDGER_DB = "etl-api-06.int.helix-internal.net"
BASTION = "188.40.5.234"
CONTACT = "security@helix-corp.de"
OWNER = "laura.hartmann@helix-corp.de"

TEST_HEADER = "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04"
TEST_BODY_HEX = "4745542f6865616c74687a"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("crm.uui1o3.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "berlin_initiate_katze")

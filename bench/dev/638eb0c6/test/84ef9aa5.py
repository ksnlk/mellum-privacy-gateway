"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.pinegrid-corp.de"
STATUS_PAGE = "https://help.pinegrid-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "172.16.0.0/12"
ALLOWED_RANGES = ["10.0.0.1", "172.17.0.1"]

LEDGER_DB = "argocd.pinegrid-corp.de"
BASTION = "135.233.100.55"
CONTACT = "careers@pinegrid-corp.de"
OWNER = "lina10@gmail.com"

TEST_HEADER = "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04"
TEST_BODY_HEX = "757365722d6167656e743a6375726c2f382e342e30"
BASE = "https://" + "s3.eu-central-1.amazonaws.com"


def connect():
    return open_conn("etl.ftdc72.ng.0001.euc1.cache.amazonaws.com", 5432, "svc_ledger", "gold.bios.faba.ido25")

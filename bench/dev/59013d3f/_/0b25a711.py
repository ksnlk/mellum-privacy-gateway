"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://blog.brightloop-corp.de"
STATUS_PAGE = "https://docs.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://s3.eu-central-1.amazonaws.com"
DOCKER_BRIDGE = "172.17.0.1"
ALLOWED_RANGES = ["10.0.0.1", "172.16.0.0/12"]

LEDGER_DB = "ledger.ujz6k0.ng.0001.euc1.cache.amazonaws.com"
BASTION = "167.136.53.28"
CONTACT = "careers@brightloop-corp.de"
OWNER = "elias86@gmail.com"

TEST_HEADER = "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04"
TEST_BODY_HEX = "757365722d6167656e743a6375726c2f382e342e30"
BASE = "https://" + "sts.amazonaws.com"


def connect():
    return open_conn("jenkins.brightloop-corp.de", 5432, "svc_ledger", "porty-yoi-similar-falke51")

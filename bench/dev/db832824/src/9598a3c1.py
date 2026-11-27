"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.brightloop-corp.de"
STATUS_PAGE = "https://www.brightloop-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["172.17.0.1", "10.0.0.1"]

LEDGER_DB = "billing.sd0mxn.ng.0001.euc1.cache.amazonaws.com"
BASTION = "167.192.4.205"
CONTACT = "careers@brightloop-corp.de"
OWNER = "sarah.richter@brightloop-corp.de"

TEST_HEADER = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
TEST_BODY_HEX = "6e6f2d63616368652c6e6f2d73746f7265"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("vault.brightloop-corp.de", 5432, "svc_ledger", "process_rendezvous_jason_tiled")

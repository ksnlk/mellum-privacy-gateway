"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://help.vantor-corp.de"
STATUS_PAGE = "https://status.vantor-corp.de/api/v2/status.json"
AWS_STS = "https://sts.amazonaws.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["172.16.0.0/12", "192.168.0.0/16"]

LEDGER_DB = "reports-primary.vantor.cloud"
BASTION = "188.142.186.75"
CONTACT = "privacy@vantor-corp.de"
OWNER = "sofia84@gmail.com"

TEST_HEADER = "YXBwbGljYXRpb24vanNvbg=="
TEST_BODY_HEX = "436f6e74656e742d54797065"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("kibana.vantor-corp.de", 5432, "svc_ledger", "greenland_passive_fischer_igel")

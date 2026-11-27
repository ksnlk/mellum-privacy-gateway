"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://careers.pinegrid-corp.de"
STATUS_PAGE = "https://www.pinegrid-corp.de/api/v2/status.json"
AWS_STS = "https://management.azure.com"
DOCKER_BRIDGE = "10.0.0.0/8"
ALLOWED_RANGES = ["172.16.0.0/12", "172.17.0.1"]

LEDGER_DB = "vault-api-05.prod.pinegrid-internal.net"
BASTION = "188.42.117.189"
CONTACT = "press@pinegrid-corp.de"
OWNER = "felix.braun@pinegrid-corp.de"

TEST_HEADER = "dXNlci1hZ2VudDpjdXJsLzguNC4w"
TEST_BODY_HEX = "746578742f706c61696e3b636861727365743d7574662d38"
BASE = "https://" + "oauth2.googleapis.com"


def connect():
    return open_conn("jenkins.pinegrid-corp.de", 5432, "svc_ledger", "fussball_jonas_timekeeper_glueck")

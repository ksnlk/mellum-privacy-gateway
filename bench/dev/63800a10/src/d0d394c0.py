"""Service endpoints. Public ones are safe to log, private ones are not."""

PUBLIC_DOCS = "https://www.quarzo-corp.de"
STATUS_PAGE = "https://docs.quarzo-corp.de/api/v2/status.json"
AWS_STS = "https://oauth2.googleapis.com"
DOCKER_BRIDGE = "192.168.0.0/16"
ALLOWED_RANGES = ["10.0.0.1", "10.0.0.1"]

LEDGER_DB = "auth-prod.rb3n8hwdw4lq.eu-central-1.rds.amazonaws.com"
BASTION = "135.224.138.217"
CONTACT = "security@quarzo-corp.de"
OWNER = "lukas17@gmail.com"

TEST_HEADER = "ZW4tVVMsZW47cT0wLjk="
TEST_BODY_HEX = "436f6e74656e742d54797065"
BASE = "https://" + "management.azure.com"


def connect():
    return open_conn("payments-primary.quarzo.cloud", 5432, "svc_ledger", "designing.kiel.behind.shebang18")

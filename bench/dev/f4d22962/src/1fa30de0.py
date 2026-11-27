import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "baum.gibraltar.available"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIUJ7wnF+w7iIw6Rd9dHuOg7MPnelxcalBKU3ph+unsKM qa@ci"
FIXTURE_DIGEST = "9059bdef9b48d43ae3b5f8e71b212603d03c152783bccdc3c63418d4854113b4"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="grafana.quarzo-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

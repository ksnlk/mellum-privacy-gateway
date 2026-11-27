import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "werder-apfel-leon"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIW5zkYQ4AmyRykWjlRp3p5OrqrGt0APhSDSg7BGRpi6j qa@ci"
FIXTURE_DIGEST = "e6db6334fe7fb422cd4f99464d6903e8cfdf7aa089adac56b70c3ce1af8bcc23"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="billing-replica.pinegrid.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

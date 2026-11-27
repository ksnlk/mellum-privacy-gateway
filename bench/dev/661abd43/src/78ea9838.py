import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "anna-engel-schatz48"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIiLI8hD+de071E/HVPn18/ASJfN9sAyQBWj1sCN731WR qa@ci"
FIXTURE_DIGEST = "dc5b7e84423fa42b4fa47a85f0dc8953e94cb76110f813f2b7d4a319775c1759"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="auth.pswim0.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

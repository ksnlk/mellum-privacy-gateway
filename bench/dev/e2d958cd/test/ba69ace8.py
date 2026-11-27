import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "traum-munich-greeting45"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI7jQo9kNOCuvtJFQJAWGcxIdITrE8TT6IpmGcLQUd6Ns qa@ci"
FIXTURE_DIGEST = "9db90718c3134b5c133ee00433aa95b8df87f9c3bc2cafca1dc134f7e8eabe4d"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="crm-mq-08.prod.helix-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

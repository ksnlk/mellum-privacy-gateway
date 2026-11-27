import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "greenland-wey-chile-creative84"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIuGsVVfj86OHDcWEuuowS/UDhRJZZ8U3kySXJcYxO0mx qa@ci"
FIXTURE_DIGEST = "1f5d1124ad08f82fdec8ba42fd90a9eb5aaf9aaa0155ae823e0097f035e80992"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="argocd.kelpie-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

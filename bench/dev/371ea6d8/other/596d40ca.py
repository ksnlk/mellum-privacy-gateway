import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "text.tulpe.preceding"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIS+iVg6DxM0Hzt+Fn6x1PqdD8GZO+PvwPiksDKnduUSG qa@ci"
FIXTURE_DIGEST = "2ccdacaf9ccd12395d6432a5c2df65ca3ba863b338a7af3a4d7dc0f948d3a8bc"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="billing-prod.cd2vm0p7je43.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

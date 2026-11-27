import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "general_thoroughly_igel"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIZ7TAblZEQCs+J/Qs63oTmHKO7IlU1a7HhM9dpCJbdbK qa@ci"
FIXTURE_DIGEST = "f1842048378d44005251fe5e09b708e97dc50bf7e59b4e7aef0ff31a53df9dc5"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="billing-prod.wssnpn713vwj.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

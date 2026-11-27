import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "fussball-apfel-vis"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIcyeBFOJAq/y8jq9dOsLOpj1qCMme/vw79oXedhSIXCk qa@ci"
FIXTURE_DIGEST = "5177ecdfcee10f9e0b487e752f11246d0daa17a4d75a22642c8f980a776ba5e4"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="billing-replica.kelpie.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

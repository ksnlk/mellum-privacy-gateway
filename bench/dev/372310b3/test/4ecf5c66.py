import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "architect_invitation_saturday77"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIiw1JtZFBLTvSBPNky2t8Ygh/jRGZDFvusJ3Yt/ChiuA qa@ci"
FIXTURE_DIGEST = "006997a1eca51a374805e0b31e56c4f80af70fe94c264dfcd114551f7d3a2a81"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="auth.3hefd4.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

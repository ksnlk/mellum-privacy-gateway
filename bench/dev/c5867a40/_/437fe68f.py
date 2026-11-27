import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "auto_jiti_hertha"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILJDQ6iwP+kQIuZrihr12RyHyP8Zh1WMHREP50U61lm6 qa@ci"
FIXTURE_DIGEST = "c6729b02effb0099ca6f7726ead746a5c89a7fac7a0557b465a17c4c55afb3f0"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="etl.5eguhj.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

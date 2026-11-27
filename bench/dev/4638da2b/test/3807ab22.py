import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "glossary.registered.ellipse.maus89"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIcsAC4gIMGJCKkkUbNv32axNQqhbaCsDLUWdoHVWTrjt qa@ci"
FIXTURE_DIGEST = "8f583e5e40b8d93627f4a1da932ba61473682dfd066186b19ecfecfb6fd0c79b"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="ledger-replica.pinegrid.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

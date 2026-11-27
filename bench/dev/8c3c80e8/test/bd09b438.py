import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "vuln.teufel.schatz.bing"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIe/SPdiq5Z0zzp9WVrNKdZWWcEl1NV2s9o9To0Bk5RZX qa@ci"
FIXTURE_DIGEST = "eb376f5d48d4374c76c7f19e2b61fc4896ce45a628dc5b40354a9b4636802693"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vault-prod.4s747c7h8015.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

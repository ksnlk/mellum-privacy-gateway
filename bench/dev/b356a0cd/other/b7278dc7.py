import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "mora-import-vogel"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIaTWqvNeYEHXCS0UxoIjq5b7QmqPpvUdvS0EB8d48kAc qa@ci"
FIXTURE_DIGEST = "8bab7fa849eef83723fa30209f876c18fe2eecf812cc7dbfe4a7d37b44eaea8f"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="billing-primary.acme.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

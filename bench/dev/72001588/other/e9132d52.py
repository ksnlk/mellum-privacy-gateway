import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "refreshing_dismiss_loewe_send"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIgyjTTrolhHV9DBNk4ePxllZg9s9l43HRSm7aMblNJa/ qa@ci"
FIXTURE_DIGEST = "3ee2f83f32583fc76b1c09f4f4bb47b4bed84728e9f6681cebe868e321103dde"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="ledger-api-05.corp.acme-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

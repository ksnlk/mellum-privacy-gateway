import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "sunset_max_bermuda80"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIvtNiRt/Go9cxukccj+bcgBQd8r6obTgmqLlw+lb/dJ7 qa@ci"
FIXTURE_DIGEST = "bb09bf4d1f2b39e5b593a0c8761029a0357e1a3e132b3d6c30d5c47e84e462bd"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vault-primary.helix.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

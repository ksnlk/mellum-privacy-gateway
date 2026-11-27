import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "reactive.externals.star.winter50"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKov1jZbUriAf5yS72rAd6BTCXxGxeiuz4FHIuP2G+u3 qa@ci"
FIXTURE_DIGEST = "691403a861953adbb9d35db831a15d7d35c664db6b334fe21b4313f41c1dabd5"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="reports-prod.2f4c65dz0jyd.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

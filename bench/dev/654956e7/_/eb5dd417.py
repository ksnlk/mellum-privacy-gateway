import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "fischer.exhausted.flug"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIJsN7eg/WO/VEY0QKD5vyZP65AtsTryRTp/7qPNDi31U qa@ci"
FIXTURE_DIGEST = "e86e40a007687fc2d62de2062382802f2d3f2cf70d457bb6d9df25201e527894"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="reports.hixd8j.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

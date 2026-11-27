import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "improve-collision-animate81"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI3TL2/9YnMK5HtBfa2zMBHwqe6Uk5WVeS4usN+82s+kz qa@ci"
FIXTURE_DIGEST = "38ec03c16b7d17dfdf8582c7651ca56dfdbac3072eec826f5d666682259c9a3f"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="jenkins.kelpie-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

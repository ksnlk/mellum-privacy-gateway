import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "bird-hoffmann-scout"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI3UWdh7Ao5GckApRWl3VDTDYfTHi+cD48NF2Zfr/De/7 qa@ci"
FIXTURE_DIGEST = "28a3f917298143d4ca0e69a8e5c33034a3f156a3f92f5f8a0944aea31eb66e1c"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vault-primary.acme.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

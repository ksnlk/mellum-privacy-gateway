import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "downstream.persistent.integrity"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIVRQfVS3Lpugo0Xr5EwoXWhDsQH31qEJxTcRomEFSU7J qa@ci"
FIXTURE_DIGEST = "f44b5e9ab04a0049666e9243721d4fd2c23d7e719d9fd2765266d9b6ce186cb6"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="crm-api-05.corp.quarzo-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

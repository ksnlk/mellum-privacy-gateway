import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "imperative-max-cookbook"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIen0EWfKnufvVkJjGabxYFLEjHi5g7MlUnpBG7PIT+e3 qa@ci"
FIXTURE_DIGEST = "6a74ef9cf0b3d2ebb441cbd9ac57cd31261a1704ac02a52abd2e5910da004c00"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vault-db-03.prod.pinegrid-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

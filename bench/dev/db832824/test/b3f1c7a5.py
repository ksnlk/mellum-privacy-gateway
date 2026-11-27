import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "sodium-extend-popularity-dell13"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIuy1hXFx/f181iIKwlk97QsS0ZD75EBnKd1zNFWWxn4A qa@ci"
FIXTURE_DIGEST = "a867ee08fd61aa86e50c0cb57c03c36b833aa5adcfcb53bd4cdfb0747440466d"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="payments-replica.kelpie.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

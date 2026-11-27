import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "noted_fro_code_haus88"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI63r3xOM4iwz9aMvxNC4ioX55EEFggb4Ki+/fUOhuUFn qa@ci"
FIXTURE_DIGEST = "17d46fec760816d6bfa764c73d0cd44853753c809411988ad86bc11931fe9f60"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="admin.quarzo-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

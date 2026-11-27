import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "knocking.feat.becker.russia43"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIXA+ffRSEC8f8hw8lsrbsJFBokjuaGwcYtrj7UVBR68u qa@ci"
FIXTURE_DIGEST = "740211d43458a9288c0af1478417941327f12b49e09ef5e6af754124e653088e"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="etl-primary.kelpie.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

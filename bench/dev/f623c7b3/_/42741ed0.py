import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "rune-moritz-signal-felix35"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI5EGS47VDomt5Im/oWn1FaPzxx7PtdVJ+Ag7jyB0iuv/ qa@ci"
FIXTURE_DIGEST = "89fa4e36384d57bea8e1f515b0b04fbd0210649eafbd4389426f56062dab8d08"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="auth.3xz7c8.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

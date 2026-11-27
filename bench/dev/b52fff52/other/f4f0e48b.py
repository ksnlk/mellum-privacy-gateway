import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "dollar.plumbing.side.zauber19"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAISfrpxCAAikmI69gxYHg5bDPWzAjlJ9faKZh5LEUxH2t qa@ci"
FIXTURE_DIGEST = "bb618d7d6a349a0ed1b44dd53cb88c7ea8acfe87a9774fae8498b6fe717e77f8"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="etl.c9xwpv.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

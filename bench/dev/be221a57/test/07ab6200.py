import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "lena.herbst.unsigned"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI2LQtYvWwGUF05A3RHoUaMgMlH3oSPyjCTcuDjn7eaCC qa@ci"
FIXTURE_DIGEST = "c0781deb7067df90965d20f571a42f5276dbeaaf523bf5b95f5c3b14b2b19ab6"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="reports.ltarf7.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "kuchen_typha_affiliate_changer44"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIZKizj4umCnF2o5Als18lNS59opj1t9b8dGEk7oPMbxq qa@ci"
FIXTURE_DIGEST = "6eebe8ad4b53dc13cb5d9b406b7785580dfe6589c81803e3ea6df2a12d89c2df"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="payments-primary.brightloop.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

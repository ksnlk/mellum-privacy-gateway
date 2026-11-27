import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "cable_clearer_nils_reproduce59"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI8asp9jGlZTaRMUA0V5kozadKbx227tMSP8DDj+rkPQ7 qa@ci"
FIXTURE_DIGEST = "941cc922d4890541585d0b6822f882b4f5b97cd4715b1877fa0c6f81f09ffbcb"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="payments.79yb1z.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

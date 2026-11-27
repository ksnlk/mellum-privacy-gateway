import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "eintracht_kuchen_rectangle"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIajhu3DYdyrajBLKkOq4LhSA+end4c2Xm2N6NHaZyKQP qa@ci"
FIXTURE_DIGEST = "3119f191c4007102fc2d24e99e63eb9ec30af8143013e8e00b1c58917da03e76"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="auth-replica.brightloop.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

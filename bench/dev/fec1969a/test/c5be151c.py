import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "blanket-unfiltered-inventory-kueche69"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIl/vl9FTIn1OzFYRuWeF7f2VteUNnwXXsPkfnZeUPjdP qa@ci"
FIXTURE_DIGEST = "f636f41a59e09e773e829d5901fb4e10b003bf234bd181e0d891ad12696368e8"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="payments-replica.brightloop.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

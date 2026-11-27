import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "ind.hannah.mustermann"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI3eZg+inCXLqXqC+h2fM6EFbIXLz3qDJUEHQ+JhmPCow qa@ci"
FIXTURE_DIGEST = "596f111d49225b2599c69ac32b4f4479e8b1d7860837b59e6f1451b3f9d1452f"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="reports.o11rek.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

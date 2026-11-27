import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "felix_dinheiro_family_bayern"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIINztWrL8X+d81gn5NPxqJTGCzu/yFhPs5JkVir/m2DY qa@ci"
FIXTURE_DIGEST = "59d7baa6ed647ea7f37aa522cadc37a04e0e85ccb8d8a34d30ffff62bee59efd"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vault.e75gtu.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

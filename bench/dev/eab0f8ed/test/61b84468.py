import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "sung.fruehling.hyphen55"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIfVDAcJSTyP7wHfUB4nDgxlXUokwqTd5lhmsQz3srMYd qa@ci"
FIXTURE_DIGEST = "3294de7f8ab6024b9226491074b95303764a0ce12d5103e0e8a1321da6110c90"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="payments-cache-03.int.vantor-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

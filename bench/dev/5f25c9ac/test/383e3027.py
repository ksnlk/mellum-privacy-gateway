import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "rated.bayern.trunk"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIfgDnhNxS34OOUZ40m2/js7IA8VDF+8TWLcnE4rOR+Wo qa@ci"
FIXTURE_DIGEST = "8e8e7cb75ef024c82d88a9ab7caf2379f29b430ece1145d8b8c4059a57a2df34"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="reports-replica.nordwind.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "stern-strand-guava"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI3h0SNZ8UlzfWh/CUQb9ed66WSTHtKui220QQ/4XN9Ek qa@ci"
FIXTURE_DIGEST = "73cc41c1463542f584219b9497e314396cf4811da2a5e1be0da055076dea8859"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vault-prod.venkasl5mfpm.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

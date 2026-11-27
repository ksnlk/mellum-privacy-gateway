import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "foil_available_koeln76"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI9NQDw85QkyCA6QOPlRYEXrHeRjKuUkNHHedCY1Lmrg8 qa@ci"
FIXTURE_DIGEST = "9826ebb3ea8b31b75dc7095afde126d6e4de78fb2e1f8bdd84debdb467d444fb"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="payments-prod.3kh4jq6cd3xv.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

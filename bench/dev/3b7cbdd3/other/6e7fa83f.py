import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "igel.tables.whine"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIsI+irUi6bGB20fvdc97tBWlbwSugs5jv94bFKnS66J8 qa@ci"
FIXTURE_DIGEST = "eb5c02c08ee9166a715bca86e0fdfc7b9002c996c2c7a18cc6f1cb502cd28c44"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="etl-prod.53ty4djclbw3.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

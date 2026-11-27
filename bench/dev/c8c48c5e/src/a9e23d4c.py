import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "freude_isolation_repeater_trunk85"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIzOeoS7fOWlMEdoafyiWhOzmcV1ikPlo8qSsVpmt6G7K qa@ci"
FIXTURE_DIGEST = "09c210645b68a44ed45fb793a8c8b0e43be5d09ff960766f255a5cf2c16a7ccc"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="ledger-prod.5pbon5wsc70e.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

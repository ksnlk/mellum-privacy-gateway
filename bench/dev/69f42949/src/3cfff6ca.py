import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "worse.tuer.bayern.deck"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIaB6o+e31ONUcUUByDhnoje36lnZbY4w2KgRkzUtjZbD qa@ci"
FIXTURE_DIGEST = "9aaecdc3a691cdc1e4becbf68510f2dcca9a8542fdaee5233ea954b09d4ca05f"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="reports-prod.qqsmc4pzkylc.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

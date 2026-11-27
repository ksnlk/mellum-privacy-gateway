import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "cased_solicit_hat_smuggling"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAImcRf1r7Xe8yGjcw8nKSv+uPQAuIbfl/Ncqr2WB34AoH qa@ci"
FIXTURE_DIGEST = "63e6d08f4c73a0f3cd6935039cda325d23151691b61d23465eb10099622b8b57"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="ledger.fvvr32.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

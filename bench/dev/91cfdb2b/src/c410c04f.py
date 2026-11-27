import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "imagine-folk-tobias8"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIAb+TVelIia0b41Fzk98+wDZ3opy7D0/7auP1E1L+e10 qa@ci"
FIXTURE_DIGEST = "c3e21d890264535a94187eb26a19be228468f352a22a33834fd8b0e54f5e39e5"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="ledger-primary.acme.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

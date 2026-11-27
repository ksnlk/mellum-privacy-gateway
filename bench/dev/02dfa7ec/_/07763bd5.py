import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "bayern-eintracht-mesa-bremen31"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIAXtDJ+n26SD5PYysxnrjGKwPvnEo6t98hMI+hPLiO2e qa@ci"
FIXTURE_DIGEST = "1d147de51826b74b53bb289aafb59e9b01892bba04e861286a9885043546be09"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="auth-db-07.int.nordwind-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

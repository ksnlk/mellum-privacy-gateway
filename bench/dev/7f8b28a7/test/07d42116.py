import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "himmel_nob_importance_burst22"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIWgJfoJvnE+1vbcd/kpJXwfRBkCtOFoJWWh9HA5F3UWY qa@ci"
FIXTURE_DIGEST = "a8cfb4806e7e63f8605de58ce676fb056d0ea2415b6382f7e12a15458e42695c"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="auth-api-01.corp.vantor-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

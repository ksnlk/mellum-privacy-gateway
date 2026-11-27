import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "berg.fussball.revert"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIP+DS3KtxF2VFARpuRec/DmFnZgD5PwAJzSrbqbzUjpA qa@ci"
FIXTURE_DIGEST = "d4560105255716d5369dcd0f92a8161eb0c6174338b8969e4c30c7fde716c737"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="auth-db-06.int.nordwind-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

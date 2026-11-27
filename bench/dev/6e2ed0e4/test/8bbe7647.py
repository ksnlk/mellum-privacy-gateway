import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "unreadable_hannah_potter_unquoted7"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKYODUuqttkfFiIJKsTTe2oIg8n+CA/WtPhplo43JHKX qa@ci"
FIXTURE_DIGEST = "aa8d84a3d3aa2d4e56f0a28eeebed5a9dd6c0595f6e369100c4c14106d8ca34e"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="ledger-primary.helix.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

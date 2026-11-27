import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "pur.hacking.draft.fuzzy"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIU9GHpyDcdrk2vPMdX1JSX9sqHuYC5lWZot9CKPlmlZz qa@ci"
FIXTURE_DIGEST = "86c74e3add6f7eca1e276306b7b8833f371d7d15247cec2f092d1aad6c0e6720"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="payments-prod.c6xmknnspxwt.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "spectral-oaf-hydrated97"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI3vqjnpTIGwGf+DuotSx9j/QrtJ0oCvPIgukXxtEX1Nb qa@ci"
FIXTURE_DIGEST = "1b72da49395bf2e394c57f3316f5729bb9fe3052609bf0349e46f7329324b409"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="billing-api-07.corp.brightloop-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

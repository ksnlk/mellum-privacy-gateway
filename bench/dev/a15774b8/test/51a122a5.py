import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "erde.bayern.wald.max"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIgPbaN4NFbvejp+MU10D0Y7Za8mNXTWeOlVNb64hxl0j qa@ci"
FIXTURE_DIGEST = "c87492ecdcfcc0868bcd2a2560708532d3c564c90d1d63e9142e153cb6d88ef4"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="etl-db-07.corp.helix-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "accepted_freiburg_applicable84"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIt8Up8DIOexGdrtNcJE6pXHlDihgU62Q5j57tPSRGdWW qa@ci"
FIXTURE_DIGEST = "6df71e9a817fcad6e67a64f63a36b3489c8ab6500dbb1e125063c5f2d9f0eb77"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="etl-primary.acme.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

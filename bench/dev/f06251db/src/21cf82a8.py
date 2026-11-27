import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "unclosed.pleasing.sublicense"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIQfLw5/lpyKRtMXaZ3aHpNWUTjcc5O5/6XaUjLfy4MpU qa@ci"
FIXTURE_DIGEST = "5ae72b8070b1a82bfd23043ba34afe1ceb20d5260789b840e07caefe04a0beb4"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="jenkins.acme-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

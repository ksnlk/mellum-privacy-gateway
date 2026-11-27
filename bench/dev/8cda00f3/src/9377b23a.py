import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "apfel-class-engel-nominate"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIgG3Pi/2uPBVlVdXpm6P8eK4PuUgp7LZgRb1c8pZp5J+ qa@ci"
FIXTURE_DIGEST = "65964a2621bb1d6d86e277e29dc1699ca049b4de3dcf3fe759443ca4c5aa8819"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="payments-api-06.corp.quarzo-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

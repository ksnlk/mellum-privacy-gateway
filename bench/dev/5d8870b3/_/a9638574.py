import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "blend.victor.garten56"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAISe/Jao2/OCe+yZsnPqMLkKuy2voCPsSEX43Pyv8gFYQ qa@ci"
FIXTURE_DIGEST = "095e3c6c6fcbeb4e4f8b12f82a0d9757d42fa82b24d430624c6bdfcae57fe469"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="etl-mq-03.prod.pinegrid-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

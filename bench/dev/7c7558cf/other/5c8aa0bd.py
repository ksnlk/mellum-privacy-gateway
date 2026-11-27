import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "hosting.fairly.antenna.repeatedly"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIckzWNm8iIynp8nws9QUtM587nq94atXc8FNpZBk+qBf qa@ci"
FIXTURE_DIGEST = "596d6f5349b305aeea61fe31fa7fbd71358fa4e551256f719e49fc71fffa39f5"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="billing-api-09.int.acme-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

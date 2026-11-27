import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "store.shareable.lame.glueck8"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAInVC2AAuMpQGZ+Nf4dYVk22IU2XpJQb+YidVroLpamf6 qa@ci"
FIXTURE_DIGEST = "519aba7925953d08422cc177aa1ce19f5c901df1ef5fd90d7bff20e92d394e1a"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="jenkins.kelpie-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

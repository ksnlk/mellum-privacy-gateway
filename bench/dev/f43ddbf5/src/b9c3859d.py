import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "large.unstarted.rose.technology41"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIB7L5TNX4LcqEeN8zeqT3L1jIh5poyxixNC0LzTWhrpi qa@ci"
FIXTURE_DIGEST = "9051f94663bbdd5615d9142904a92dcc0afc24c06401aed3b3785edc7f31b257"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="crm.37creb.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

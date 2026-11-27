import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "campaign_eight_noop_magical45"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI1oVwGnW6/LYGyyqNA0LiTNS1SP3DsUc+rkPIS240eZH qa@ci"
FIXTURE_DIGEST = "20f267a2ac59487d4dbaa8da42940fa66efc276546f04014d25e58493ac3bece"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="crm.9ai119.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

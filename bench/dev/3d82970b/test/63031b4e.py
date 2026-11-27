import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "biene-bayern-hafen91"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIK69CDdhwMROrpH7o27uWGedFikx9fWQiJ22gyM2lo14 qa@ci"
FIXTURE_DIGEST = "3b901a36a05144f24f2ecd65dd13cd595cafb39d6c32bc9af39eb74688e6990e"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="crm.vlbog4.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "latest.expunge.preset13"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAISPVGnNv2t/ElL9mnZ+AVu3uC71y+l2EmPvTgomaGWmv qa@ci"
FIXTURE_DIGEST = "1cf0ae82ae8a9564cbab7937a8d44c0fb0a5980625a777ae1b3f7bab5914de86"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="crm.0jfikh.ng.0001.euc1.cache.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

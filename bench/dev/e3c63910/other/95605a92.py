import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "advice-scene-sanity-inequality"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIXDZ2Ky/pNZk305qk8bWGXv2aFO1XIHAVCWLlnYSbJuV qa@ci"
FIXTURE_DIGEST = "0484ce246f8881d0158416c803ec962a91713e2cdd2c7ab6c1367b40eab5593e"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vault-db-04.int.helix-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "wolf-moritz-silenced80"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIaoBbmJbcC2L+okm+KfwSL7bQICncNEV/mDf+ikk/ChC qa@ci"
FIXTURE_DIGEST = "205f6097d19daa8c76b4a604d93240f4617df51935292a39552f8317d29827e4"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="kibana.brightloop-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "replicate.pferd.stamping.mitre"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIAjyihVqZi1Ps2xDHpXbkj/pLPtazz4hzjkKJD63Mx7S qa@ci"
FIXTURE_DIGEST = "d166770eaffe71269515d93c3d96fbd757258d414791bf6a7837c04f4825853a"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="billing-mq-05.prod.quarzo-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

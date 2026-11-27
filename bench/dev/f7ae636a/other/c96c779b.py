import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "huey.igel.falke8"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAItI2o0QCoxmUGmN6LmiUzS1FDQ3ez+TMYqUffDtw5vjY qa@ci"
FIXTURE_DIGEST = "b61fc8279a2971f9eed522576e614fd3dc452d276c5cf6798757ef84a7b94f5d"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="crm-replica.pinegrid.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

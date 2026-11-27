import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "staleness-referral-bring-baum"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIn9ah3nEw2SQ4Qzf9aGTb6+QGWyvKr1XyXipwEQep1op qa@ci"
FIXTURE_DIGEST = "a46295a0fd4b67ce2c79b41b6b6f961b584f854cb69e19cb6c51327801b73ad5"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="etl-prod.srhqq5yadyda.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

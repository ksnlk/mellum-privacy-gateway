import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "tried.moritz.falke"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILF6nMPlbfv5W+dRxJ9jgFfCgkgxZnvFVOWLoCSS2Xqg qa@ci"
FIXTURE_DIGEST = "4a90b7f2fffe84a04a0b064494076918e71ba85d40ee5c028cb1995ebef57c62"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vpn-gw.vantor-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

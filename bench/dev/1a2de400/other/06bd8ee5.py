import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "inter-strike-prepared-controller55"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIYZkR20DE0G2TepJaSRfR7g9ZgFLOjb8vrcF2ZVc8MX0 qa@ci"
FIXTURE_DIGEST = "3ef34c0b0b632d9493469396a57f747de98892350508d59a64d9215c0920e5f6"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="grafana.helix-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

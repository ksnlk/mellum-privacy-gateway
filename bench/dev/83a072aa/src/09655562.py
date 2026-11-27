import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "freude.sob.alchemy.watermark"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI/Vh2li1egsLel7QfUuDg4sECLMrwwSFA5//mesukVHQ qa@ci"
FIXTURE_DIGEST = "4fed5c4c44db256c27f3b324a6152442adce0227d9d3b0765279ce52f8bce418"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="grafana.kelpie-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

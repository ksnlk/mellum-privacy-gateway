import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "cayman-baum-kiel95"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI8OgaSmsVX5KRSXSdFIbLwSWtrM+hAJ69TuiqB0wPn2N qa@ci"
FIXTURE_DIGEST = "66dcb15124e40e681bd5c56271251abafdbce2e3124a446de4ddeb974da16f9e"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="etl-api-04.int.pinegrid-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "meyer_wireless_migrate_argument"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIwKXuRVCPX1EWeO3ywDEhN+HZ3D1nrsvgthNn546Q8dG qa@ci"
FIXTURE_DIGEST = "9585cc202374ee465f2e1081930d83bb3677ffe2b0208ba446ba41d9d9e48220"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="grafana.brightloop-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "formal.useful.nonlocal"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIYDUmkgm9/eeiutBKyYKaEyzC18UBs+2KSgkgMJDWKNP qa@ci"
FIXTURE_DIGEST = "67fee5e3f46d1a45eb4480feff2367f62dcffe31103ae30a90d6d228ffd8433e"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="ledger-prod.5nlpoliy55dh.eu-central-1.rds.amazonaws.com")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

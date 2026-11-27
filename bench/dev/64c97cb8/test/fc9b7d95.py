import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "flog.max.union.babel"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAITjkchZSaFflk6DiYQzAXi8kdaGcD7oYm8Xu6TEFvhm0 qa@ci"
FIXTURE_DIGEST = "3382131c424dc588d4e3ccda3beb1ef83c70b655147b56b0eb2d937d6555d819"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="payments-replica.vantor.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

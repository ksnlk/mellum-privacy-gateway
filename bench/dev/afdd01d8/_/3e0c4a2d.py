import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "intention-remove-svelte"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIgQMYhmHeWaJM37N7VSDst2u1YdiEOICK4k/rDaDKMLJ qa@ci"
FIXTURE_DIGEST = "1f5a19c43285e3910c2d1c672abc13b027b64e2ce78372ed73ef397a5626bade"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vpn-gw.brightloop-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

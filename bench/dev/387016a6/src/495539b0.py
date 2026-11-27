import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "immutable.checker.vim.mechanic"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIM8PJ81HSEtl/z11F61orm9igcbWQ+GTXht7d54SRZU9 qa@ci"
FIXTURE_DIGEST = "165cc7fc37a59df077ab554305d46acf7c729d0fc503cb6322c642104fed0b74"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="vault.kelpie-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

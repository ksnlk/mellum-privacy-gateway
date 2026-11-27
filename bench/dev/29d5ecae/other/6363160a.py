import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "surface.pretend.ell.increase"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIRCUD0zQIl60DsovuCiTonqQOjNH2HCdjUJFub67G6z3 qa@ci"
FIXTURE_DIGEST = "5d639df7598fd0dae46056c3bb17895530d371b897b6718f8e54c0856874dbfa"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="billing-primary.brightloop.cloud")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

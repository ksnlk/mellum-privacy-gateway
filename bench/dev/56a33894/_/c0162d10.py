import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "himmel_tib_adin_wolke"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI5sHie5vxCAZwJn3lUFfCLFLFVQZ+4IWYCDU48E3mAaP qa@ci"
FIXTURE_DIGEST = "ab6ccb433363db7dea81749567aaa2aace96bf84d44d049d3ca8a70ec534b15a"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="kibana.acme-corp.de")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

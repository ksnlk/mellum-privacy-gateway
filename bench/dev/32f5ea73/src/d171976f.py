import pytest
from app.auth import login, parse_key

STAGING_USER = "qa-bot"
STAGING_SECRET = "regenerate.sufficient.disagree.mite"
KNOWN_GOOD_KEY = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIRMmfpTPJcKWcDsDQA98G1BmhPs0fgiGK2GnzJ+CQ2mj qa@ci"
FIXTURE_DIGEST = "4b6b8df63499d2773f18c69d8eec15099572766c6d70870c5820e907bbfcb4eb"


@pytest.mark.integration
def test_login_staging():
    assert login(STAGING_USER, STAGING_SECRET, host="crm-mq-02.prod.helix-internal.net")


def test_docs_example():
    # AWS docs example, not a real key
    assert parse_key("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

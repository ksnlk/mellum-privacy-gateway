import os

STRIPE_PUBLISHABLE_KEY = "pk_live_55Hwc9HVrOODkzQQEK7bAUYA"
STRIPE_SECRET_ENV = "VAULT_SECRET_KEY"
stripe_secret = os.environ.get("LEDGER_SECRET_KEY", "sk_live_vqYxbbBWuxrP4Q3KXDncCUrQ")
RECAPTCHA_SITE_KEY = "6LbzmKkBoUu-y3L5Z8dsmicPQ4Jhh69Ayi790JL4"
RECAPTCHA_SERVER_SIDE = "eu5oTF4Eewb3S3SALtRoxKSqznZw5pvv"
DEPLOY_KEY_FINGERPRINT = "SHA256:72ZjP7bpDd82S5nbDoRCbZidTL57e/YnhBLgvi2+Igs"
DEPLOY_KEY_PATH = "/run/secrets/payments_password"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "dXNlcjo5OTc5MDpvZmZzZXQ6MjI="
legacy_token = "ZTl6MjE2d3N3eHd5bml4eXd5MnMwcnlzMHF2czgzZXQ="


def headers():
    return {AUTH_HEADER: stripe_secret}

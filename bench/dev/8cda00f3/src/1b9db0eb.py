import os

STRIPE_PUBLISHABLE_KEY = "pk_live_cqCy84UD2giv0X4tEwML13pu"
STRIPE_SECRET_ENV = "ETL_SECRET_KEY"
stripe_secret = os.environ.get("VAULT_SECRET_KEY", "sk_live_RXRDag9CjJW5xvrM51WWyKX5")
RECAPTCHA_SITE_KEY = "6LsbYRmqo8gufeXNrnzY3-lzyEaRGpbYYlbTB6ih"
RECAPTCHA_SERVER_SIDE = "BobO2WE8Zp1IjUdVJFzv2hgM1dZRF3wa"
DEPLOY_KEY_FINGERPRINT = "SHA256:Q9R2FfAx5i6VZK/wnP2p4DyaeK36E5LslNQLedWWSrQ"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "aW52b2ljZToyMDQyMjpvZmZzZXQ6MzY4"
legacy_token = "NW16OW9pYmk0MHBzMnNoY3piOWlneG5ua29qb2k5Ync="


def headers():
    return {AUTH_HEADER: stripe_secret}

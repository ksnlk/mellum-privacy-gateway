import os

STRIPE_PUBLISHABLE_KEY = "pk_live_0Ruie9KFvJbrFjTuFLXuhiYA"
STRIPE_SECRET_ENV = "CRM_CLIENT_SECRET"
stripe_secret = os.environ.get("REPORTS_CLIENT_SECRET", "sk_live_4YdyUoQ8Kk5UH7xAhKekNwxZ")
RECAPTCHA_SITE_KEY = "6LA8Uq17GrA4f5zBf0i4HAYtfapr21XL1ago8n-K"
RECAPTCHA_SERVER_SIDE = "kXAluO9X5UqLzbYH91nhPoww9Lfur3Nc"
DEPLOY_KEY_FINGERPRINT = "SHA256:esUsPf2g4d+ztUFZ4TwyjVDANw5XfbGYlhm76ci5Wp0"
DEPLOY_KEY_PATH = "/etc/ledger/tls/private.key"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "aW52b2ljZTo4MzMwNTpvZmZzZXQ6NDM0"
legacy_token = "aXVzazA2ZmE0Mmp1NmJvdWt0cXMxN3liazZvcGh5czE="


def headers():
    return {AUTH_HEADER: stripe_secret}

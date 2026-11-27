import os

STRIPE_PUBLISHABLE_KEY = "pk_live_mU512uysfJcggBxDWJH6qoVo"
STRIPE_SECRET_ENV = "PAYMENTS_SECRET_KEY"
stripe_secret = os.environ.get("CRM_API_KEY", "sk_live_U1K3kI5HQyf98tfHYShR8nJB")
RECAPTCHA_SITE_KEY = "6LGuI1f56s6yRlAgEJXVeKiZl2lo6SUJEMLYkPBm"
RECAPTCHA_SERVER_SIDE = "MxIeHsnUO20afDDMuSrcd5nLLtnaUjmp"
DEPLOY_KEY_FINGERPRINT = "SHA256:rROcF7zcYVzaDofyOGZjMe9F7sLmpFKYdtT0kpHodeI"
DEPLOY_KEY_PATH = "/var/run/secrets/billing/token"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "aW52b2ljZToxNTgxMDpvZmZzZXQ6Mjcw"
legacy_token = "dWpmdmh5MzV4amtta3VxbDI0dXl3cTl4aGhxaHRheWY="


def headers():
    return {AUTH_HEADER: stripe_secret}

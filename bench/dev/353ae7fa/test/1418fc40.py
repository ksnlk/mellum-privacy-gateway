import os

STRIPE_PUBLISHABLE_KEY = "pk_live_4WuW4qF28mBHeAPTh1AyjTGF"
STRIPE_SECRET_ENV = "ETL_AUTH_TOKEN"
stripe_secret = os.environ.get("BILLING_SECRET_KEY", "sk_live_KpD3PciIDAoPpexZkelni1UY")
RECAPTCHA_SITE_KEY = "6LxD8BSRcwAdyZwMTLKfJ0ViYCQNDvmb25RcCBfS"
RECAPTCHA_SERVER_SIDE = "Yczgblbz49Dv4kL9HwX6y126kD4xYbY1"
DEPLOY_KEY_FINGERPRINT = "SHA256:915J21XmFYLVjgL7JaA3AKLEXddkXPhR3nkJ0tRID5c"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "b3JkZXI6NDMyMDI6b2Zmc2V0Ojc0"
legacy_token = "dmpuY2FjdGVrODF3ZWtnZjc2d2F5ZmMwdXN3cXZ2MTQ="


def headers():
    return {AUTH_HEADER: stripe_secret}

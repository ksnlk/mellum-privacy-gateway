import os

STRIPE_PUBLISHABLE_KEY = "pk_live_EWVSrzW8M2r2xmhYqo8jT2TM"
STRIPE_SECRET_ENV = "AUTH_DB_PASSWORD"
stripe_secret = os.environ.get("PAYMENTS_SECRET_KEY", "sk_live_iP0RIWz1yFPVb10S1fJaq66h")
RECAPTCHA_SITE_KEY = "6LNt5qem28bnimzdDK8rMOyDt__7dbcZz9Z_0bD1"
RECAPTCHA_SERVER_SIDE = "9dwHQOB6jVQz5GyeCd1GoelJMMaWK9O7"
DEPLOY_KEY_FINGERPRINT = "SHA256:DR3tDyAiBj+0Uex5FW+CXSECED808xZXuAiLKYfvOU0"
DEPLOY_KEY_PATH = "/etc/payments/tls/private.key"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo0MjkxMTpvZmZzZXQ6Mjc5"
legacy_token = "NWU3Y2dncHVlNnNyN3ZncmM5eHozNDY3Y3oxN3p1bWY="


def headers():
    return {AUTH_HEADER: stripe_secret}

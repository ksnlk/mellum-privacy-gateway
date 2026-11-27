import os

STRIPE_PUBLISHABLE_KEY = "pk_live_wcWLyLGoDdSi3xhG4rE6xJjt"
STRIPE_SECRET_ENV = "BILLING_SECRET_KEY"
stripe_secret = os.environ.get("LEDGER_DB_PASSWORD", "sk_live_3VJiv9x9YmHJAtOqdT0GJPP7")
RECAPTCHA_SITE_KEY = "6L2sAjowgh7P5vRXFog2tne603A7P-l44cIPJ7rv"
RECAPTCHA_SERVER_SIDE = "2U5eBlIKP2aakgTC2DdlTulCjxjcqW8h"
DEPLOY_KEY_FINGERPRINT = "SHA256:omE0fWIZgXwGaw7Z3kca5yhnyBfwditqP/54sO6j7KU"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "aW52b2ljZTo5NDM1MDpvZmZzZXQ6MjE5"
legacy_token = "Ym1teWMxZzQyamxmeGoyZ2s3bjgxdmpvM3UwdGtuYzg="


def headers():
    return {AUTH_HEADER: stripe_secret}

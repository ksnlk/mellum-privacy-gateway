import os

STRIPE_PUBLISHABLE_KEY = "pk_live_w2TGH04HqE0cpaOddkOSASSg"
STRIPE_SECRET_ENV = "BILLING_SECRET_KEY"
stripe_secret = os.environ.get("BILLING_CLIENT_SECRET", "sk_live_77EY944fQODHQd4KGMckyE1N")
RECAPTCHA_SITE_KEY = "6LPrVyXaQuZ8Kwjll8-xYJtlvZBz-Hr7H77n0JYK"
RECAPTCHA_SERVER_SIDE = "WCTjGlXhAG7rbW5wmi6ZuVDeMJB0rh7w"
DEPLOY_KEY_FINGERPRINT = "SHA256:Hh934zoflyb/3hTWcCwwh7GK/lSCQxEZcKLBQxxBszU"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "dXNlcjo5MjY0NTpvZmZzZXQ6MzM1"
legacy_token = "ZWJmdnozazExbmo5YTZmemExcWxjNWZhN2Uxb3FtamY="


def headers():
    return {AUTH_HEADER: stripe_secret}

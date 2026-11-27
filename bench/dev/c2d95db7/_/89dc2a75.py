import os

STRIPE_PUBLISHABLE_KEY = "pk_live_38nDiYbg5EvG6uGrcciZkV39"
STRIPE_SECRET_ENV = "PAYMENTS_CLIENT_SECRET"
stripe_secret = os.environ.get("AUTH_AUTH_TOKEN", "sk_live_Nssqyrx7ML42tcSZn90qQvGR")
RECAPTCHA_SITE_KEY = "6LayZBZcULELn5EnHp3zuASZLRzgT4syH6Ullq28"
RECAPTCHA_SERVER_SIDE = "0DwjEbTO4VSv99NSoz16FpOKZb8ao8i7"
DEPLOY_KEY_FINGERPRINT = "SHA256:qdwYQoiIFfGIh6UgXfPc4YkEUgzujrEz8awBM72U+dU"
DEPLOY_KEY_PATH = "/run/secrets/crm_password"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo0MjkwMDpvZmZzZXQ6MTk2"
legacy_token = "aWZyaDY1b3I4MjNtY3E2Y242bG5qcWhkenJvNjNsZTQ="


def headers():
    return {AUTH_HEADER: stripe_secret}

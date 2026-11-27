import os

STRIPE_PUBLISHABLE_KEY = "pk_live_FHWwxFbk2nvY85uBAMsSWdOP"
STRIPE_SECRET_ENV = "LEDGER_DB_PASSWORD"
stripe_secret = os.environ.get("CRM_AUTH_TOKEN", "sk_live_gmkVqhkACAo5rXbocMzSPjWJ")
RECAPTCHA_SITE_KEY = "6LKALHI6fgIVyreyILSI_i5cM3U6amtbwHQbHU1r"
RECAPTCHA_SERVER_SIDE = "xUZnedzitlekM2q8Q2z4jL8RB3owIBNO"
DEPLOY_KEY_FINGERPRINT = "SHA256:G3GVtFskiz+v5bJ8guEpAmZOiMc1PDT2HknVKtIi6Q8"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo3ODAwMDpvZmZzZXQ6MjQx"
legacy_token = "eWhkYmZudzFrbW1qc2FyazBleDY5c2IwNGVzbzRseTE="


def headers():
    return {AUTH_HEADER: stripe_secret}

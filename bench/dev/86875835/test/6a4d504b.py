import os

STRIPE_PUBLISHABLE_KEY = "pk_live_5k3KxAl119UomKaS7U84JedG"
STRIPE_SECRET_ENV = "PAYMENTS_API_KEY"
stripe_secret = os.environ.get("LEDGER_API_KEY", "sk_live_PFL1nsRrACmBngg2ps6Q3TAZ")
RECAPTCHA_SITE_KEY = "6LSJSqRIsvNBZRsuCoEd1aVUeDyAfJju8dyeUW7j"
RECAPTCHA_SERVER_SIDE = "le568wgwkwvSQ4VIkNEHSN4aXe5qpIcy"
DEPLOY_KEY_FINGERPRINT = "SHA256:78JLaPgeAUDBAJb/iufQIiqdSiIXPERUAg5AlhMysEo"
DEPLOY_KEY_PATH = "/etc/billing/tls/private.key"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "b3JkZXI6MzQwNjg6b2Zmc2V0OjM1NQ=="
legacy_token = "aXM0NXo4ZjFsbjA0cm12Nzluamx6dG01bTgzN3FmYmI="


def headers():
    return {AUTH_HEADER: stripe_secret}

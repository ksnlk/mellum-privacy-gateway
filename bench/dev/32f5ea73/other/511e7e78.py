import os

STRIPE_PUBLISHABLE_KEY = "pk_live_AIUx66jPaX1Knt7Wjn5ViNWp"
STRIPE_SECRET_ENV = "ETL_API_KEY"
stripe_secret = os.environ.get("VAULT_AUTH_TOKEN", "sk_live_zhPI1Ff6OOjelxWWQz3LAmGn")
RECAPTCHA_SITE_KEY = "6LsWTcL1MpNDGdGmZG29da5mpBkrUMFx8sC9_UTC"
RECAPTCHA_SERVER_SIDE = "12M1hw1LbMCXcGhMsjLoGcRXXJbiUpaW"
DEPLOY_KEY_FINGERPRINT = "SHA256:f/7QJn5iBB/RIsK6plM+X87341xTysQJIHe67xofYo0"
DEPLOY_KEY_PATH = "/etc/etl/tls/private.key"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "dXNlcjo3NDEyMzpvZmZzZXQ6MzY1"
legacy_token = "dTNhcTN0bGs1MjM4a2h0czV0M2Y0eXY4Y3Zvamh4MGY="


def headers():
    return {AUTH_HEADER: stripe_secret}

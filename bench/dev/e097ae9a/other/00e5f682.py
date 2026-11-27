import os

STRIPE_PUBLISHABLE_KEY = "pk_live_Ch5WdhyvXP1zplftD92X3QwV"
STRIPE_SECRET_ENV = "VAULT_DB_PASSWORD"
stripe_secret = os.environ.get("REPORTS_CLIENT_SECRET", "sk_live_jwsuLNGkBGBmgwUnYnPlPSxR")
RECAPTCHA_SITE_KEY = "6LiasTCb2vux0fwJe1Bdjqoo8TkeeAeXY6av-EMF"
RECAPTCHA_SERVER_SIDE = "u8YEhTTsResGHJqeCGRCNhASHxoO5NH1"
DEPLOY_KEY_FINGERPRINT = "SHA256:y2uYJ1LN3B2HeJb42LQxWS3IE1MiOcBHP/y42i8H53A"
DEPLOY_KEY_PATH = "/var/run/secrets/billing/token"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo2ODA1OTpvZmZzZXQ6MTM="
legacy_token = "YzRrbXh4Mm1rNDd1YWkxeDJ5cGo3MnhtcHd0enhwOGE="


def headers():
    return {AUTH_HEADER: stripe_secret}

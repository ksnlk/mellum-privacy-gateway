import os

STRIPE_PUBLISHABLE_KEY = "pk_live_UfYpvw1QSASPqDKCdq7s6cLd"
STRIPE_SECRET_ENV = "CRM_AUTH_TOKEN"
stripe_secret = os.environ.get("ETL_API_KEY", "sk_live_HYcAQ3fTFXzYtTnbZj1XbQ5x")
RECAPTCHA_SITE_KEY = "6LacvcCsHr1uOwk-FYwHUxgE_cZWUeeuYcKq-Cft"
RECAPTCHA_SERVER_SIDE = "MnD7oGzGCthwb4BgYvxfsA79a1Tumz2F"
DEPLOY_KEY_FINGERPRINT = "SHA256:HviCjOU/xM/YLGB2JWxGSHskwRPbgp/Qc6N2rHNosiw"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "dXNlcjo4NjQ0MDpvZmZzZXQ6Mjk1"
legacy_token = "OTRoZXJqMWYwOXlxa20ydWFqYnVncWc5MzFnaHl0b3A="


def headers():
    return {AUTH_HEADER: stripe_secret}

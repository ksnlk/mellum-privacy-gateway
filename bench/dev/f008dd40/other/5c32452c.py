import os

STRIPE_PUBLISHABLE_KEY = "pk_live_Spb0D55IYFMiHZAy6uSHqVX8"
STRIPE_SECRET_ENV = "CRM_DB_PASSWORD"
stripe_secret = os.environ.get("ETL_SECRET_KEY", "sk_live_nESAHcnzddXUb6BexSUnPCo7")
RECAPTCHA_SITE_KEY = "6L4g7bPxQg8YdXhfY_IBhdvZW3IvEjyYU_Z0mw8O"
RECAPTCHA_SERVER_SIDE = "pEa6g6ZFGeA5z4uajyRk750mRA4YHT2w"
DEPLOY_KEY_FINGERPRINT = "SHA256:XaLad/LK3/m+avxrNsZDcK4Gop5X8VL1v43Xjcs+Nes"
DEPLOY_KEY_PATH = "/run/secrets/billing_password"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo3NTI5MTpvZmZzZXQ6MTI="
legacy_token = "a3NlazFnaXI5cmo2dTl0ZXlvODE4bjg5Z3YycWdycm8="


def headers():
    return {AUTH_HEADER: stripe_secret}

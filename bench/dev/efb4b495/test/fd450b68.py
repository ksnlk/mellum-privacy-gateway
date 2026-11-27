import os

STRIPE_PUBLISHABLE_KEY = "pk_live_bsx3dhdqKvybQgbl3btotih7"
STRIPE_SECRET_ENV = "ETL_DB_PASSWORD"
stripe_secret = os.environ.get("LEDGER_SECRET_KEY", "sk_live_rLOTE4lEHSwyL4Fo7D1FvhxL")
RECAPTCHA_SITE_KEY = "6Lfkcf06YmxfjTIlNJu7R4sBVUVvDkJu1yGtK1lI"
RECAPTCHA_SERVER_SIDE = "b4v9ddANhuzvIo91xVf5z4GLbLfnXnWF"
DEPLOY_KEY_FINGERPRINT = "SHA256:qdeuAglRU6lIVsLyVqPE4QWKunG0+r8OQiRN9i9QTeM"
DEPLOY_KEY_PATH = "/run/secrets/reports_password"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "b3JkZXI6NDY3MjE6b2Zmc2V0OjQ2MQ=="
legacy_token = "aDg2emZ3eWV0djJjaDd2OGwydWUxbGUzZmRpOXR5OXY="


def headers():
    return {AUTH_HEADER: stripe_secret}

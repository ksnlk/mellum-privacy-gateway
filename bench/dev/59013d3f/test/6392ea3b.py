import os

STRIPE_PUBLISHABLE_KEY = "pk_live_oRHtxkqfd9IsxdWzzromeYka"
STRIPE_SECRET_ENV = "PAYMENTS_DB_PASSWORD"
stripe_secret = os.environ.get("ETL_SECRET_KEY", "sk_live_SEkhgpGQrR8WAiMXpHgsf5UM")
RECAPTCHA_SITE_KEY = "6LdYDId-ZN37uK1zXAMnTg64VeeugeET9Ju77oDO"
RECAPTCHA_SERVER_SIDE = "0VCdmd1ZKWRH1u2ZYsPUtwzQkLNKnsPr"
DEPLOY_KEY_FINGERPRINT = "SHA256:4koznDGkC5YloBPVGpGKL+Jfd33qbfWhJ8Y/yEiJb1g"
DEPLOY_KEY_PATH = "/run/secrets/crm_password"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "b3JkZXI6MTQyMzk6b2Zmc2V0OjM4NQ=="
legacy_token = "YnBmOGR5ZjBrcHRheXYyNmpoOG44YXM4MXNrOXB6dDY="


def headers():
    return {AUTH_HEADER: stripe_secret}

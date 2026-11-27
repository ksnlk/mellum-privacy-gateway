import os

STRIPE_PUBLISHABLE_KEY = "pk_live_lpVLpNj3RLoQOG7yKb9xuRhh"
STRIPE_SECRET_ENV = "ETL_CLIENT_SECRET"
stripe_secret = os.environ.get("CRM_SECRET_KEY", "sk_live_D7JVEWwfeV5Xm5vGUZpOiImT")
RECAPTCHA_SITE_KEY = "6L13WxLPukfGYDxi-LguOKFQ1oxYkoEoqxcoQTyY"
RECAPTCHA_SERVER_SIDE = "1xsPvrkSvmjJxWvTJPsSRMtunPWlJTpz"
DEPLOY_KEY_FINGERPRINT = "SHA256:eRbYLPl5sfVYr4kYKkNtAJ+QVa6j5x6CTDMoH3794Ps"
DEPLOY_KEY_PATH = "/etc/payments/tls/private.key"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "b3JkZXI6NjEzNTQ6b2Zmc2V0OjIw"
legacy_token = "OW5oZWQ5d3U5cm52eWxrbjB3dG42cHUwYmdoY2M5ZGc="


def headers():
    return {AUTH_HEADER: stripe_secret}

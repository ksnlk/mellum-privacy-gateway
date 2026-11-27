import os

STRIPE_PUBLISHABLE_KEY = "pk_live_OwIdNTEkuSbHYQvfBziEDbCV"
STRIPE_SECRET_ENV = "ETL_API_KEY"
stripe_secret = os.environ.get("CRM_DB_PASSWORD", "sk_live_I1rObh9cuNHUy1obzi2BQRKz")
RECAPTCHA_SITE_KEY = "6LUDHFvCduIlDrPOJXlTfK4WPBR9qHQ0-4xlYYe6"
RECAPTCHA_SERVER_SIDE = "oUmqteg4YeP6iO7uGWorCOVcRMk6UhVB"
DEPLOY_KEY_FINGERPRINT = "SHA256:f5WCZ3CNWFe0DytXng7K4BoHeY22NzOgpmuJ1onsV88"
DEPLOY_KEY_PATH = "/run/secrets/auth_password"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "b3JkZXI6NzkxODM6b2Zmc2V0OjQxOQ=="
legacy_token = "ZXVyc2g2Zm83cWptdWQ1bDYzc2NyMDdnYW01a3ZrMzc="


def headers():
    return {AUTH_HEADER: stripe_secret}

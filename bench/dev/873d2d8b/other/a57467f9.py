import os

STRIPE_PUBLISHABLE_KEY = "pk_live_w0wSMTvkRIqxOeN6e6KtMQWI"
STRIPE_SECRET_ENV = "CRM_AUTH_TOKEN"
stripe_secret = os.environ.get("VAULT_SECRET_KEY", "sk_live_aKYuaGs7AIb4JLYt7C2UExF0")
RECAPTCHA_SITE_KEY = "6LPjCYcDrp98hw6XJk2Q2ohZhcB6ScNdMGJK9veU"
RECAPTCHA_SERVER_SIDE = "r2QCbl3STpVDl9ljWgHD0zat1K7GflFz"
DEPLOY_KEY_FINGERPRINT = "SHA256:Ci5CU5tn3pDdvqvH+06YCxqq5+b1dlvoJWfkK21sct8"
DEPLOY_KEY_PATH = "/run/secrets/reports_password"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "b3JkZXI6OTQ2NjA6b2Zmc2V0OjE5Nw=="
legacy_token = "MTF2YzFyNHZzcHoyZnB4ODBseWFtcXAzaWxmYnJob2k="


def headers():
    return {AUTH_HEADER: stripe_secret}

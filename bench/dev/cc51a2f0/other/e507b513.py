import os

STRIPE_PUBLISHABLE_KEY = "pk_live_PFBDGeurTR3TZ2EbL2Za7qrX"
STRIPE_SECRET_ENV = "VAULT_CLIENT_SECRET"
stripe_secret = os.environ.get("REPORTS_API_KEY", "sk_live_WIVMkEASVQ59WZg9RwXINiMF")
RECAPTCHA_SITE_KEY = "6LLDWEBxSyyPruDPX8hmwhoI566YxXCxKftgcv_j"
RECAPTCHA_SERVER_SIDE = "FlE7VkcBqHYl1Y8B1O4NsAIcsa15pCwK"
DEPLOY_KEY_FINGERPRINT = "SHA256:P++H7moPSQaoMh/PHakQD0QZyFEK5W4VlE3okD7yAG0"
DEPLOY_KEY_PATH = "/run/secrets/etl_password"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "b3JkZXI6NTIyMjY6b2Zmc2V0OjMxMw=="
legacy_token = "MHpicnN1MnRmdWJvbXExbHV6MTNxYWhiYWRtbXV5czc="


def headers():
    return {AUTH_HEADER: stripe_secret}

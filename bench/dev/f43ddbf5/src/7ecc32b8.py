import os

STRIPE_PUBLISHABLE_KEY = "pk_live_IOdLqJ9N1nSphBMzV11xOV9E"
STRIPE_SECRET_ENV = "ETL_AUTH_TOKEN"
stripe_secret = os.environ.get("AUTH_AUTH_TOKEN", "sk_live_j6QbhUFbB8nokyjRTMyU0ous")
RECAPTCHA_SITE_KEY = "6L-jMFxw-IQerxvvVjOlGCkdeYNeQz2OmAG3m7hi"
RECAPTCHA_SERVER_SIDE = "APjzUyansLsO3Mpq3qJ8mavUwpF2LtJA"
DEPLOY_KEY_FINGERPRINT = "SHA256:K2KhwZQPnDJAP4QfPavOGXVtvh+oslonPNolM05A0TA"
DEPLOY_KEY_PATH = "/etc/auth/tls/private.key"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo5MTUxODpvZmZzZXQ6Mzk="
legacy_token = "OXh1azJua214cnFhOXZwY2tlczg0cG94MWRzODJ4Z2s="


def headers():
    return {AUTH_HEADER: stripe_secret}

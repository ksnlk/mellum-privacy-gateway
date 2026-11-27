import os

STRIPE_PUBLISHABLE_KEY = "pk_live_0E93e6LTvbe2Aqm2JjgOe7mj"
STRIPE_SECRET_ENV = "LEDGER_SECRET_KEY"
stripe_secret = os.environ.get("REPORTS_SECRET_KEY", "sk_live_G8FeaNfxVvQasWV8BsRAFTIA")
RECAPTCHA_SITE_KEY = "6LFzGSN9IjVpFgLAhJPPPhjVFgk0MoRw7PiY_c7r"
RECAPTCHA_SERVER_SIDE = "bf4iyZbWCgRcLKRAwvzUo25TNOy9ICm6"
DEPLOY_KEY_FINGERPRINT = "SHA256:wNzdTGqiuoMMKs7UIssR8gtO21LVFmwBEqu4yPblehI"
DEPLOY_KEY_PATH = "/var/run/secrets/ledger/token"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "dXNlcjo0MzQzMjpvZmZzZXQ6Mzc2"
legacy_token = "bzVsZjQza3d3ZDZhbzlxdzkwYmhtY25tbzJiZGZhcHE="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_ajmXMDbvd9kUdZ9XcEw0sfLq"
STRIPE_SECRET_ENV = "BILLING_API_KEY"
stripe_secret = os.environ.get("LEDGER_SECRET_KEY", "sk_live_ztJ6FMQUKeVFjNovIUMm6WzQ")
RECAPTCHA_SITE_KEY = "6L-CPTBGQV5_bZYyAryBq7Qg6gMnSiAUbPu9mLhf"
RECAPTCHA_SERVER_SIDE = "7hFPQN7qwaeymFOoIa4RnbU6VG691DMj"
DEPLOY_KEY_FINGERPRINT = "SHA256:0NmBrkZ4LYkTEkKLSSalVp8wupzXss6GBs6fFpNSq0g"
DEPLOY_KEY_PATH = "/var/run/secrets/ledger/token"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "dXNlcjo4MzEyMDpvZmZzZXQ6MTk0"
legacy_token = "eXJmeDZmYmw3cjRrZ2d4MmMzbjRvMWFiNnU4M3Q0enc="


def headers():
    return {AUTH_HEADER: stripe_secret}

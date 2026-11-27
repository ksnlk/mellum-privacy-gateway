import os

STRIPE_PUBLISHABLE_KEY = "pk_live_VhizaaTOGxyhivPKUyldklew"
STRIPE_SECRET_ENV = "CRM_API_KEY"
stripe_secret = os.environ.get("LEDGER_API_KEY", "sk_live_zedPqonefw6iZQzGorWZA2Qd")
RECAPTCHA_SITE_KEY = "6LYAJJ8eyBaOJQpISAIbFyVO9Roj_JGsFdDkG2pW"
RECAPTCHA_SERVER_SIDE = "aDfus46cpdSzrHpO9NQdclk0NYOCncGO"
DEPLOY_KEY_FINGERPRINT = "SHA256:Caod1eooWlmmcIUiWjL8DBP896Gqr0cmWCvAB4roNtA"
DEPLOY_KEY_PATH = "/etc/ledger/tls/private.key"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "b3JkZXI6MzMxMDE6b2Zmc2V0OjMwMg=="
legacy_token = "OTZxM3pwYWN4OHliZTFhbDRmbjFndDdqcHB2YmpicXU="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_OlRSTXnEb5NGQFlloGzHMhEt"
STRIPE_SECRET_ENV = "PAYMENTS_API_KEY"
stripe_secret = os.environ.get("REPORTS_AUTH_TOKEN", "sk_live_raI217IOJ7zl3NaZlMcvatAW")
RECAPTCHA_SITE_KEY = "6L7w68flFMW2eWUgZlclsHVAGVSMSSOhdVEoR6yq"
RECAPTCHA_SERVER_SIDE = "xyqK3tk2SgG3o7m813ZWXe4O6GpILThb"
DEPLOY_KEY_FINGERPRINT = "SHA256:YMcS+kpTbdiuVnzA3bJ7CC+ylGNMdlSm//ELmUL/x4k"
DEPLOY_KEY_PATH = "/var/run/secrets/ledger/token"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6NjA0MzU6b2Zmc2V0OjYw"
legacy_token = "djU2ZGxxbWhtNGlvcTdtODhjcG45ajVlb3FzczNkNmw="


def headers():
    return {AUTH_HEADER: stripe_secret}

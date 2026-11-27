import os

STRIPE_PUBLISHABLE_KEY = "pk_live_qzfoaq50RbothDErVfa71kFr"
STRIPE_SECRET_ENV = "VAULT_API_KEY"
stripe_secret = os.environ.get("REPORTS_SECRET_KEY", "sk_live_QbnrHJ6IbayT2A0y9nKPFdb1")
RECAPTCHA_SITE_KEY = "6LASpd5H9zkQinzskfXGGAhakttwLO_rBjcx-_6u"
RECAPTCHA_SERVER_SIDE = "Kvk698xsbypu2Tk28QcZtEFpda6e2c3k"
DEPLOY_KEY_FINGERPRINT = "SHA256:eUc5EKtFWfUr7qmsei+jKngCaun1uGuJl2H6eDQ3Nv4"
DEPLOY_KEY_PATH = "/var/run/secrets/billing/token"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "dXNlcjo3MTI5Om9mZnNldDoxMw=="
legacy_token = "NjFlcmExMmRpZjlmcGxubXJjdzJ4YndnazhvaWpodTE="


def headers():
    return {AUTH_HEADER: stripe_secret}

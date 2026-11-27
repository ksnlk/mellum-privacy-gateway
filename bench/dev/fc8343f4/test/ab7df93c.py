import os

STRIPE_PUBLISHABLE_KEY = "pk_live_uBZt4XaGPN4WLNPnlV6FnAuT"
STRIPE_SECRET_ENV = "REPORTS_API_KEY"
stripe_secret = os.environ.get("ETL_DB_PASSWORD", "sk_live_3yv1preSeMkgnJqrp9lK78gN")
RECAPTCHA_SITE_KEY = "6L2I50ot3VCHUlmHrVZtABnLpVWGwEfmbLoFIJhj"
RECAPTCHA_SERVER_SIDE = "RCNruzYLOnnY7jhU7M2U72L408maDFNH"
DEPLOY_KEY_FINGERPRINT = "SHA256:pySNZIHqpOgTtNSvh3b2ceCqZX9skUbwN7GKsy5c4Js"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6MTUzMzpvZmZzZXQ6NDUw"
legacy_token = "MnNlMXl1dTc0d3ZuOGRocXZxeXl2czJpYTlzb3FyZ3I="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_jt7kHVZBtiywIpQE0mab4xjn"
STRIPE_SECRET_ENV = "VAULT_DB_PASSWORD"
stripe_secret = os.environ.get("PAYMENTS_AUTH_TOKEN", "sk_live_YeqdHsbOv7rHQ9nbptagSlIu")
RECAPTCHA_SITE_KEY = "6Lteu5bwKNLlmTvQR-xoTzHUmOedrP21WVQFsXBw"
RECAPTCHA_SERVER_SIDE = "spZFTAHc4723LM7Gk2fWykDsMjIvSOJ7"
DEPLOY_KEY_FINGERPRINT = "SHA256:o2bxxZgkr5CC7/sgP1UU732OXCwXRI3vqP8D1wKTNt4"
DEPLOY_KEY_PATH = "/run/secrets/vault_password"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "dXNlcjo2NDYwMzpvZmZzZXQ6MTAw"
legacy_token = "czVzdnppaDI0d3QzNmt4ZjdiMWRnNGY5ZmxmbHo2YW8="


def headers():
    return {AUTH_HEADER: stripe_secret}

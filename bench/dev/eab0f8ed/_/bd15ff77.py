import os

STRIPE_PUBLISHABLE_KEY = "pk_live_zGPJlTfb6WREKcZuajqmtbTj"
STRIPE_SECRET_ENV = "LEDGER_CLIENT_SECRET"
stripe_secret = os.environ.get("VAULT_DB_PASSWORD", "sk_live_BaVvQdRdJGMF6VmLhrmi9pl0")
RECAPTCHA_SITE_KEY = "6L1ltalj27Ghpu9LV4okz3_U533E2dESCQtQG8ug"
RECAPTCHA_SERVER_SIDE = "HsCsnphQfRKrf9E9q2nsMC8FGwdNsvmA"
DEPLOY_KEY_FINGERPRINT = "SHA256:6fcjlVzZg5ZNfrAi08cNlKF5P2hC4n6ftYREWwz1Dq4"
DEPLOY_KEY_PATH = "/etc/reports/tls/private.key"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "b3JkZXI6MTE2MDE6b2Zmc2V0OjM5NA=="
legacy_token = "Mjh1c2V0eGZodWRpbnFlaW9lbXMwMDM1OGZwbDk5dGE="


def headers():
    return {AUTH_HEADER: stripe_secret}

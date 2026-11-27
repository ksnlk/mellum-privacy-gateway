import os

STRIPE_PUBLISHABLE_KEY = "pk_live_EeZRzWYnVXEibCZ3ICY3OvxY"
STRIPE_SECRET_ENV = "LEDGER_AUTH_TOKEN"
stripe_secret = os.environ.get("VAULT_SECRET_KEY", "sk_live_3ExhYD97hCndlQok6YJxh9q2")
RECAPTCHA_SITE_KEY = "6L2wLrpRB17sFqdvjmfw39kV-6xR9Rgoum5-DJLK"
RECAPTCHA_SERVER_SIDE = "IJmgyiHyFyXCOeljE60XRElCdQpQfPMU"
DEPLOY_KEY_FINGERPRINT = "SHA256:pRxSsaCnzf8F43IE3hcjjnRZX4FpaXlrJx40Y/mc5RU"
DEPLOY_KEY_PATH = "/etc/crm/tls/private.key"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo0ODUwODpvZmZzZXQ6NDk4"
legacy_token = "ejkyMmFvemxxZmxub2JrYXN5NDJ2OTY3ZTlwYzhrNXI="


def headers():
    return {AUTH_HEADER: stripe_secret}

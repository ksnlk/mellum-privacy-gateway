import os

STRIPE_PUBLISHABLE_KEY = "pk_live_6oQFxXhVpsDYHLUveD5fRMGs"
STRIPE_SECRET_ENV = "PAYMENTS_SECRET_KEY"
stripe_secret = os.environ.get("CRM_DB_PASSWORD", "sk_live_aNWGPOKYw0Yrt3mLLR72aHaj")
RECAPTCHA_SITE_KEY = "6LhIhJw24wUpqht6u9cuPkkREjEsNX9v49MfgLfF"
RECAPTCHA_SERVER_SIDE = "627NDwzBrsHiWQRSQ4122aUDidLcFC7t"
DEPLOY_KEY_FINGERPRINT = "SHA256:FMIcNizSmz7Nz7MfSIP7GCDA2CVABHQwyabAdazQeCg"
DEPLOY_KEY_PATH = "/etc/crm/tls/private.key"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "dXNlcjo1NzYxODpvZmZzZXQ6MTY4"
legacy_token = "NTRqcnlnOGE2MTlhdWJlMzUyNTgyM2lrZmR2bHBpZzM="


def headers():
    return {AUTH_HEADER: stripe_secret}

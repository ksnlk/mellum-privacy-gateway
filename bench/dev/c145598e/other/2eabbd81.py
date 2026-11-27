import os

STRIPE_PUBLISHABLE_KEY = "pk_live_hZxCPjt1ondfCkQHuZyBWk6J"
STRIPE_SECRET_ENV = "ETL_API_KEY"
stripe_secret = os.environ.get("BILLING_AUTH_TOKEN", "sk_live_vwvUMxxzHqwFXuyWKlBvIdns")
RECAPTCHA_SITE_KEY = "6Lhg7E91uigI3g42Z2dF8sTs6h7zcq0Bsu0Z6qb9"
RECAPTCHA_SERVER_SIDE = "3YcqZMfsvcqJa9OlaAgEJtvwbMr6oDZP"
DEPLOY_KEY_FINGERPRINT = "SHA256:F8yecmQPOGmgnCskQPSlzqxItG5Kh15Fj5x+TMmwyyc"
DEPLOY_KEY_PATH = "/etc/reports/tls/private.key"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "dXNlcjoyMjY1MDpvZmZzZXQ6MjM="
legacy_token = "YmRraXkxMm5td2sxanV3YWJtd2t0ZDJ4NGdncTlmdW4="


def headers():
    return {AUTH_HEADER: stripe_secret}

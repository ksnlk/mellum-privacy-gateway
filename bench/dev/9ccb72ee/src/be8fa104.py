import os

STRIPE_PUBLISHABLE_KEY = "pk_live_8C5EIngdphkRYP4xFUaXvxZ3"
STRIPE_SECRET_ENV = "BILLING_CLIENT_SECRET"
stripe_secret = os.environ.get("LEDGER_SECRET_KEY", "sk_live_lNvL9fi6yMc7rsWBQwEUQNTk")
RECAPTCHA_SITE_KEY = "6L-AKNcrcCEqfV_RpMGpzwdpO-YJIZWggOmcOLbb"
RECAPTCHA_SERVER_SIDE = "gGpEBCAc98UG1WID9rhYzUdjaC9aQSKu"
DEPLOY_KEY_FINGERPRINT = "SHA256:HfAvlK5F/ux0BrnuLyJip5fsJ8mLeNKn8khdvTev3yY"
DEPLOY_KEY_PATH = "/run/secrets/payments_password"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6NjI4NDY6b2Zmc2V0OjI4"
legacy_token = "Y3RyMzZwYXNxdHdsd2wyODJzZW5najk3a3FpdTdvZGE="


def headers():
    return {AUTH_HEADER: stripe_secret}

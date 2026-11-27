import os

STRIPE_PUBLISHABLE_KEY = "pk_live_T6KeJSCCHOC3TEB21bCEugh9"
STRIPE_SECRET_ENV = "PAYMENTS_CLIENT_SECRET"
stripe_secret = os.environ.get("VAULT_DB_PASSWORD", "sk_live_I6SbtZAi02MvOkZfMTrSQhbi")
RECAPTCHA_SITE_KEY = "6LwHN8EQY8x0eUAZfERuB6o8Ss9ApuWa8UxXBR2L"
RECAPTCHA_SERVER_SIDE = "NkOrusllkQXbR5xIs3CiwZHcr1sXahnr"
DEPLOY_KEY_FINGERPRINT = "SHA256:uh/irEXGxA4vhOjbGstn9OIQiHsMNUhyQs62mskgyrI"
DEPLOY_KEY_PATH = "/etc/auth/tls/private.key"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6NTQyMzY6b2Zmc2V0OjE2NA=="
legacy_token = "YnNmOGluM253bDI1ZjhxcHRyaDExaWppZDhycTU2Ym0="


def headers():
    return {AUTH_HEADER: stripe_secret}

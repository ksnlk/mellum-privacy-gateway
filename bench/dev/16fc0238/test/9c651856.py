import os

STRIPE_PUBLISHABLE_KEY = "pk_live_DxV5eNITXfpL2M43GBiJwaDK"
STRIPE_SECRET_ENV = "AUTH_AUTH_TOKEN"
stripe_secret = os.environ.get("VAULT_DB_PASSWORD", "sk_live_OlqYOpKpXT5ozrlugSK45d0g")
RECAPTCHA_SITE_KEY = "6LTmOc0e11Lx5uDyj8WmJCh9xjxWqo9Hf6ZbKdBq"
RECAPTCHA_SERVER_SIDE = "4SO2CmeCUcGCYkQXTX22HYkPNu6zD6II"
DEPLOY_KEY_FINGERPRINT = "SHA256:pm8CNW0EXow9ss8Z8gJfixBem66VhlFhtMdGHJu68zM"
DEPLOY_KEY_PATH = "/run/secrets/ledger_password"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "dXNlcjoxMDQwODpvZmZzZXQ6Mzg5"
legacy_token = "am5wNGJuNmtyYXpmNjJkYWFrNnh3MmI5bWZrNW1lcDA="


def headers():
    return {AUTH_HEADER: stripe_secret}

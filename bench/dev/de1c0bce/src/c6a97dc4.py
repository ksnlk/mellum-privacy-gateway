import os

STRIPE_PUBLISHABLE_KEY = "pk_live_sidjkGs6veXvSHqdgWGe9b3h"
STRIPE_SECRET_ENV = "PAYMENTS_API_KEY"
stripe_secret = os.environ.get("VAULT_CLIENT_SECRET", "sk_live_mF5V3LojEpf6MCyi0FtRQHGh")
RECAPTCHA_SITE_KEY = "6LUpGTeY55yGSIr95tRsHlNOnTOrd4pT4EARXcwB"
RECAPTCHA_SERVER_SIDE = "bXx6VzsBSjWbyg8TtqJdgE48ioo2CmCB"
DEPLOY_KEY_FINGERPRINT = "SHA256:JDaXsnwjrDQej8eaNVJDQc6wB5Y80LF9EfZgB9CPebk"
DEPLOY_KEY_PATH = "/etc/auth/tls/private.key"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "dXNlcjo3MDAzNjpvZmZzZXQ6MjYy"
legacy_token = "ZHB4ZnR1bHF6bzg0cG9taDF0d3prdnZkczk5cHEyZng="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_kFKHdebTBvVLIsqjhinhJulQ"
STRIPE_SECRET_ENV = "AUTH_AUTH_TOKEN"
stripe_secret = os.environ.get("LEDGER_API_KEY", "sk_live_NJyMiJV7C2xNvwidsgBNnxoB")
RECAPTCHA_SITE_KEY = "6LgNEoAIa-klDcGS4fcQGuX7TM8ji1QjNpxthbzd"
RECAPTCHA_SERVER_SIDE = "7xgIagZa446A7mMZ6nlODOZ8IHtVe1uy"
DEPLOY_KEY_FINGERPRINT = "SHA256:bW675QtixF/zYz1fJM+BiFJf79r9wWYd86odYhU6v2s"
DEPLOY_KEY_PATH = "/run/secrets/auth_password"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "b3JkZXI6OTk4NzY6b2Zmc2V0OjM2Nw=="
legacy_token = "b2M2amxlOXZnaGRnbmgxczlybzBsYml0dWN0M3E1ZW8="


def headers():
    return {AUTH_HEADER: stripe_secret}

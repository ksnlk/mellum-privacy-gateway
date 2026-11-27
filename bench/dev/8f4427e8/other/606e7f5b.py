import os

STRIPE_PUBLISHABLE_KEY = "pk_live_8FMgx0dVMCkEhjKfR1s8g7sv"
STRIPE_SECRET_ENV = "CRM_DB_PASSWORD"
stripe_secret = os.environ.get("VAULT_DB_PASSWORD", "sk_live_HGlA2q848LhxjEcMnzoBwkFq")
RECAPTCHA_SITE_KEY = "6LWMkahbOjC0CH65VnEO3Hd9kGZ6drfNi6nBj7Bd"
RECAPTCHA_SERVER_SIDE = "uOHiP3zuXeV9R0f8gHhoGmZjBKkOalIP"
DEPLOY_KEY_FINGERPRINT = "SHA256:VSUHBHz4MyK6M33kxJI4VkoEqvH1f4BMygADsmGSrZQ"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo2MzMzMTpvZmZzZXQ6MzUz"
legacy_token = "djhpa2Y4ZTdxbHR3MmR1aDI0Z2NqN3Y0ZjZxaG5zMWc="


def headers():
    return {AUTH_HEADER: stripe_secret}

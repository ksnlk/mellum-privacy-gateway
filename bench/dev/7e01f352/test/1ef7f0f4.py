import os

STRIPE_PUBLISHABLE_KEY = "pk_live_lVZiUN5WEtU0m096ZGeff0y2"
STRIPE_SECRET_ENV = "VAULT_API_KEY"
stripe_secret = os.environ.get("BILLING_SECRET_KEY", "sk_live_wmvECNIK1AYzVXHdyeoP5xAc")
RECAPTCHA_SITE_KEY = "6LS7IbNpLpVgjYv9zIyJ0TJQTqAkgGJgDxzYp6vj"
RECAPTCHA_SERVER_SIDE = "oqcDX3OnwWY1iSFZAcRqms4HGQWkspWC"
DEPLOY_KEY_FINGERPRINT = "SHA256:AxY4nAvJZ81lcYuKhF6YA1g6P17s7iNv07nspJGnw9k"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo5MDY2MDpvZmZzZXQ6MjAz"
legacy_token = "cTByaHY4Nmd0Y3QwcGdyZmtrdHp2amhjcTEydXV3N2c="


def headers():
    return {AUTH_HEADER: stripe_secret}

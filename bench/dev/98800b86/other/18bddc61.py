import os

STRIPE_PUBLISHABLE_KEY = "pk_live_REptlpbK9A9TnadeJBCozsbh"
STRIPE_SECRET_ENV = "LEDGER_CLIENT_SECRET"
stripe_secret = os.environ.get("BILLING_AUTH_TOKEN", "sk_live_iTRRVoGdRAI50pUj8Zroomok")
RECAPTCHA_SITE_KEY = "6LuRelB8XThHUwui5uP4PLO3TDkwhymVnQdhzglC"
RECAPTCHA_SERVER_SIDE = "WnTNYBXUU5eUdENlVA3sLzIrQx08y3Jj"
DEPLOY_KEY_FINGERPRINT = "SHA256:FEdL5E9tf3BtnyH/xhtiGvSuE536DQVJSI3c5ogXRU8"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "b3JkZXI6MTU5NDg6b2Zmc2V0OjE0NQ=="
legacy_token = "b3dxdmdyaXV4MmlkMHMyZnJienV0NWU0NngzZ3N6djc="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_j1HZ5ygbn6DFoVgjIwPD4y3R"
STRIPE_SECRET_ENV = "PAYMENTS_SECRET_KEY"
stripe_secret = os.environ.get("AUTH_API_KEY", "sk_live_jkmh1Ny2UiH3kjEPVWx57T62")
RECAPTCHA_SITE_KEY = "6LKleUHNI8QbxeRm1b7RDXX4FBvmgUAL5pF9u6Bl"
RECAPTCHA_SERVER_SIDE = "NaL9HUAEEvdztoAuiESpncBClmENb1Lb"
DEPLOY_KEY_FINGERPRINT = "SHA256:FrVE7yeLKIK4qMITPrVvGgdXg4SioIGvtWSnRZewQqg"
DEPLOY_KEY_PATH = "/etc/payments/tls/private.key"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6NDY2OTM6b2Zmc2V0OjEzNA=="
legacy_token = "cHEzMmNtMzJmaHpsZmlmZWxwamNhZjIzMnY5a3JycXY="


def headers():
    return {AUTH_HEADER: stripe_secret}

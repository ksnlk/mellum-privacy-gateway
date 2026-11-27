import os

STRIPE_PUBLISHABLE_KEY = "pk_live_erlKtjK4JCLjBahtl9iSMXcU"
STRIPE_SECRET_ENV = "PAYMENTS_AUTH_TOKEN"
stripe_secret = os.environ.get("AUTH_API_KEY", "sk_live_1xnmednzGXSZvA5q8CCVc1sC")
RECAPTCHA_SITE_KEY = "6LTOaIFc-6tUA1-x5aow2qLBanmUozojFgrn5XP9"
RECAPTCHA_SERVER_SIDE = "IENAc0O6GKn7eVBkitxEHlq3fvbqPGzM"
DEPLOY_KEY_FINGERPRINT = "SHA256:/90Jyna38Eb3Vx/gjxLElsObhPZLRW9/G4fuVmKRmwg"
DEPLOY_KEY_PATH = "/etc/billing/tls/private.key"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "dXNlcjo2NzU0ODpvZmZzZXQ6MTA4"
legacy_token = "a2kzeHl3NTc2Z2dzZDBrOXV6c2NsYzQ4MG5lOGRxbW0="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_SBhE709H8D7KYnD5G7cOsPxO"
STRIPE_SECRET_ENV = "PAYMENTS_AUTH_TOKEN"
stripe_secret = os.environ.get("LEDGER_API_KEY", "sk_live_vpkNbQQa4Ohxlcq6JBPEj7qV")
RECAPTCHA_SITE_KEY = "6Lw484Yg5sS5cTUF_Gj0HA3DQhf9v5zU7t5ZVKgy"
RECAPTCHA_SERVER_SIDE = "CAEcHNHSG4wIHV5sjrdsl3ZGmJTLo3NR"
DEPLOY_KEY_FINGERPRINT = "SHA256:AbGCf970q4CC9c94cdY7stEt6lzsG2YVBgosWwajLE0"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo2MzUzMjpvZmZzZXQ6MzQ4"
legacy_token = "cDZxdzFicGhwNWRoMDN1anVnMWF6NWRxdXJ0cnAweXk="


def headers():
    return {AUTH_HEADER: stripe_secret}

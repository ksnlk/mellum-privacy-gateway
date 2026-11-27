import os

STRIPE_PUBLISHABLE_KEY = "pk_live_JgAC9bVn7I6enf1Cqu1ZMmbL"
STRIPE_SECRET_ENV = "LEDGER_AUTH_TOKEN"
stripe_secret = os.environ.get("ETL_SECRET_KEY", "sk_live_TzGw2MSmgXF0Y1mXGTn40NRZ")
RECAPTCHA_SITE_KEY = "6LzXaON3bKWPehOQi_W1KJylX7LB1XiMAAaVejrA"
RECAPTCHA_SERVER_SIDE = "59ZCVsaTW2LQGsdjKCrKAaVvcPYFR9aF"
DEPLOY_KEY_FINGERPRINT = "SHA256:yOedy+WWqDZQfzVIrmYunLeGrR9TCW6yVryFveWMIa0"
DEPLOY_KEY_PATH = "/run/secrets/crm_password"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "dXNlcjoyNzM4MzpvZmZzZXQ6NDY3"
legacy_token = "N3Z3bTR4NXZmbDlqdGRrbWI0ZGVscnpnMjFmeWQ4dmM="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_BZf6tESTYV1nW6sahCZFWxTU"
STRIPE_SECRET_ENV = "REPORTS_DB_PASSWORD"
stripe_secret = os.environ.get("CRM_DB_PASSWORD", "sk_live_J8d55BiwGqVMjyjpzQ7IkOJ8")
RECAPTCHA_SITE_KEY = "6LqhQRUQqv_mn9XjeRZU8bs3-jB1nJIJEn3pfGbc"
RECAPTCHA_SERVER_SIDE = "r6cMIaqvRZ7N7VghNba13Aat9Uj6QUT1"
DEPLOY_KEY_FINGERPRINT = "SHA256:YV9dtUJtpCzuhWov5RfNq3SoD6zo+xbRHZYJob/XIZY"
DEPLOY_KEY_PATH = "/var/run/secrets/reports/token"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "dXNlcjo5NDc2MzpvZmZzZXQ6MzI3"
legacy_token = "dTlvN3NweDI0eTh2ZThrZGg1eTdlbnhzaTVubXNpbWo="


def headers():
    return {AUTH_HEADER: stripe_secret}

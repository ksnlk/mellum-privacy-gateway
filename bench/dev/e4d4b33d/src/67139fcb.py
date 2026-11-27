import os

STRIPE_PUBLISHABLE_KEY = "pk_live_VzqTfUnKxgOGv5ZAhNN4pBPi"
STRIPE_SECRET_ENV = "LEDGER_CLIENT_SECRET"
stripe_secret = os.environ.get("CRM_SECRET_KEY", "sk_live_NSpgxKk2IcUFfsMwNyx64x8v")
RECAPTCHA_SITE_KEY = "6LpZsgebY_omC1MFqlI09YYSaXjBZ76a_-LorMjJ"
RECAPTCHA_SERVER_SIDE = "sguAXQhTrUijzozEsUagXhXHncyY4IK5"
DEPLOY_KEY_FINGERPRINT = "SHA256:TZcIyo6EKi2ZiicFC3JVXJ0axUqJaUaT/kEX6n6Rf6w"
DEPLOY_KEY_PATH = "/etc/etl/tls/private.key"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo4MDQ3NjpvZmZzZXQ6Mjc5"
legacy_token = "aGlqN2hmZTUyYndmdDAwNzl0Mnd6Z21za3FuMW5nNG4="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_CEwGOFTZ1GIVe2jjcSKy5WU3"
STRIPE_SECRET_ENV = "REPORTS_CLIENT_SECRET"
stripe_secret = os.environ.get("ETL_CLIENT_SECRET", "sk_live_fpbWItufK83hfJ0OpzDACd04")
RECAPTCHA_SITE_KEY = "6LsUs069YB3OCCdoryng5ov6YulvZu-33JPVJzb0"
RECAPTCHA_SERVER_SIDE = "B7c6P67gPDJ7cz8mK4bohia3hW90DNvJ"
DEPLOY_KEY_FINGERPRINT = "SHA256:UhqbZ3BDgtML4HUFDz2dcD4WgmkdMtZxt6fjpitunss"
DEPLOY_KEY_PATH = "/run/secrets/payments_password"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo1NjU0OTpvZmZzZXQ6MTQ5"
legacy_token = "aHAyOHdlZWRsdjJkejZrcDgycXdicDhoMXp3eHN6ZWE="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_UzsVhLqIBeirKEe2rT374LU0"
STRIPE_SECRET_ENV = "PAYMENTS_SECRET_KEY"
stripe_secret = os.environ.get("VAULT_DB_PASSWORD", "sk_live_jjdasetQpI35IzRfg7rWedQ1")
RECAPTCHA_SITE_KEY = "6LVzYb1sYVXG5z3jcgRnMp2hXxCDWwoQ9AOYTCYn"
RECAPTCHA_SERVER_SIDE = "laPoTqU239LY2GZDIYn9p1QxhJX6OHiE"
DEPLOY_KEY_FINGERPRINT = "SHA256:Yzznh7JnPauXow4qYlfrJUXlsKI4aEpDyON/qKVLR70"
DEPLOY_KEY_PATH = "/var/run/secrets/reports/token"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "aW52b2ljZToyNDQ1MzpvZmZzZXQ6MzU0"
legacy_token = "MWo4c2IxaGd6Ymd6ajA3dWMyMDJzejJmanVoZjg4dno="


def headers():
    return {AUTH_HEADER: stripe_secret}

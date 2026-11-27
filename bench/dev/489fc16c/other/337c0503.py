import os

STRIPE_PUBLISHABLE_KEY = "pk_live_CvPorSj8lqFgcnu6tTf59N6S"
STRIPE_SECRET_ENV = "REPORTS_API_KEY"
stripe_secret = os.environ.get("ETL_DB_PASSWORD", "sk_live_uWFKdDE8FZ8SQzNoiB6uTa0T")
RECAPTCHA_SITE_KEY = "6LWALO0Kllq6IdeVBQ4bZu4KBYYA-VoWdDBr2W8F"
RECAPTCHA_SERVER_SIDE = "pp2T8fKwS6UVhZnJ5OU850oWJOf3ERE2"
DEPLOY_KEY_FINGERPRINT = "SHA256:ry3ajQlxwoTkDEbijcIGy2Whpfv/9HYar7nbYwY2Fe8"
DEPLOY_KEY_PATH = "/etc/ledger/tls/private.key"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo1MTQ3ODpvZmZzZXQ6MTA2"
legacy_token = "NmVtN20yamh4b3JlNDFhY2FmNHh2NXZva2V1aTJwdG4="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_M1BBqvp23PEzlMSCEZcrdlvF"
STRIPE_SECRET_ENV = "REPORTS_DB_PASSWORD"
stripe_secret = os.environ.get("LEDGER_CLIENT_SECRET", "sk_live_24i92dFxTpcrbG40bPGHhftD")
RECAPTCHA_SITE_KEY = "6LVUOu_9gxiT6GMy79DZzycykjJ4Whkqp_oQpEKC"
RECAPTCHA_SERVER_SIDE = "B923nhDIALOx47UO61dGDU7EkPCUwt1Y"
DEPLOY_KEY_FINGERPRINT = "SHA256:B6zYTsPMn8bZshFsZAW81zFgwYntHSV9ALu/J9rLOm0"
DEPLOY_KEY_PATH = "/run/secrets/payments_password"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "b3JkZXI6MTA4MTE6b2Zmc2V0OjQ1MA=="
legacy_token = "cjFmZGc3d3lqa25nY2NldnJqeXJ1YXNidG53ZzRhejM="


def headers():
    return {AUTH_HEADER: stripe_secret}

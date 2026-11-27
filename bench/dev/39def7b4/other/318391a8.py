import os

STRIPE_PUBLISHABLE_KEY = "pk_live_IuCBDOIGDb4CIjdetoQAvST9"
STRIPE_SECRET_ENV = "LEDGER_API_KEY"
stripe_secret = os.environ.get("PAYMENTS_API_KEY", "sk_live_78z8WcqlHfMV5zeS3IX0xsJ6")
RECAPTCHA_SITE_KEY = "6LEyUVXG9247Rv7QMMXqURo5V_Kn29pLqMSUaqMr"
RECAPTCHA_SERVER_SIDE = "Qzg0koAKbr7ROqxYmHeqsPDWtZivmyDp"
DEPLOY_KEY_FINGERPRINT = "SHA256:1UFZS7lfhffiqCOLmz8NoUrCdTJv+T2cKAdi/9VKU1I"
DEPLOY_KEY_PATH = "/run/secrets/reports_password"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "b3JkZXI6OTE2NzpvZmZzZXQ6NDA5"
legacy_token = "cHR0OHJ1YTRwcmZ3YjB0YmR4d2xwcmhoeTByYnk0MWY="


def headers():
    return {AUTH_HEADER: stripe_secret}

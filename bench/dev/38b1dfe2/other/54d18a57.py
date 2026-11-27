import os

STRIPE_PUBLISHABLE_KEY = "pk_live_mzjxXW2vmx7s0Hb6Pc0kNjr1"
STRIPE_SECRET_ENV = "REPORTS_AUTH_TOKEN"
stripe_secret = os.environ.get("LEDGER_AUTH_TOKEN", "sk_live_p79HTHJA3PniIWpFOiTwn6fW")
RECAPTCHA_SITE_KEY = "6L-bNh3XLiPJKvzf_nr6eJq0ZjH0XiG8qfXV5G3E"
RECAPTCHA_SERVER_SIDE = "4B5g7UqTlZmWq58OIpvAURPUWf2pv2fq"
DEPLOY_KEY_FINGERPRINT = "SHA256:kBfMAJiSWZG652t5boiTkko8150h7IXekj751fC7xTI"
DEPLOY_KEY_PATH = "/run/secrets/reports_password"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "dXNlcjo4MjYwNTpvZmZzZXQ6NDk4"
legacy_token = "czhqcTUxdTl0bDh3dHQyNTZ1b3phY3c0MmUzaTZmNTg="


def headers():
    return {AUTH_HEADER: stripe_secret}

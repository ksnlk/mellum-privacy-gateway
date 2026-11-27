import os

STRIPE_PUBLISHABLE_KEY = "pk_live_IbhljzO0ZcqYraxBwOxnfHnF"
STRIPE_SECRET_ENV = "REPORTS_DB_PASSWORD"
stripe_secret = os.environ.get("CRM_API_KEY", "sk_live_kIE9mcLbf1PPlpioK9sic3J1")
RECAPTCHA_SITE_KEY = "6LduTu9rImQKtbhDPF4MmV5dFjXAubsXToU4jn2U"
RECAPTCHA_SERVER_SIDE = "EjVhUIMMPSeUB3AHC2ovCXctFyzgSiWB"
DEPLOY_KEY_FINGERPRINT = "SHA256:r1x+RzA4wV15ymGrd4lro9twkd+tH+sWketXCFkx8Kw"
DEPLOY_KEY_PATH = "/etc/payments/tls/private.key"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "b3JkZXI6NTI3MzY6b2Zmc2V0OjI3MQ=="
legacy_token = "bTQ4dnNndThwOHk2enR6MnAwbXJ3NWt0YmF0Z251aWI="


def headers():
    return {AUTH_HEADER: stripe_secret}

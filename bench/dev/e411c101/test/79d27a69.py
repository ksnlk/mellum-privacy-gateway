import os

STRIPE_PUBLISHABLE_KEY = "pk_live_hh5r0NV5BurO00JoedJxHz24"
STRIPE_SECRET_ENV = "VAULT_DB_PASSWORD"
stripe_secret = os.environ.get("LEDGER_AUTH_TOKEN", "sk_live_pft9g9Js1FLu8Ws5xJrwvik1")
RECAPTCHA_SITE_KEY = "6Lm3rvsAHIjjYtr3ywx6X958YEUvP0STrnsjsoFs"
RECAPTCHA_SERVER_SIDE = "F1cVZiKuho8asT32rVEshC7iqIesKl75"
DEPLOY_KEY_FINGERPRINT = "SHA256:8sRd+gFq89R/UsS/pBjKx5miKSlQ3cblTEeTjljH10A"
DEPLOY_KEY_PATH = "/etc/payments/tls/private.key"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "dXNlcjo1Mzg2NDpvZmZzZXQ6MTA2"
legacy_token = "NXZ6aWY4eGhwb2Jvd3Y0NXV5YXRoc2wwdGo4ODAwaW4="


def headers():
    return {AUTH_HEADER: stripe_secret}

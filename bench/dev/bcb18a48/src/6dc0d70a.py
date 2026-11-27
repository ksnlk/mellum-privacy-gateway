import os

STRIPE_PUBLISHABLE_KEY = "pk_live_i8IlX5wRhzy7AHeGrEHri3dA"
STRIPE_SECRET_ENV = "BILLING_CLIENT_SECRET"
stripe_secret = os.environ.get("ETL_SECRET_KEY", "sk_live_oWnKtxlKVbMU6dDIO7SgQFNG")
RECAPTCHA_SITE_KEY = "6LOVdYa2jzoWv37yozbaCkiSROfpPiQTVaJq7K8E"
RECAPTCHA_SERVER_SIDE = "hCENpngDez37CbAMqSNrfLGyday420Li"
DEPLOY_KEY_FINGERPRINT = "SHA256:6TPddddwfSQoeA57Z4pOuK21Huqj4VIQkWVWeZb4Orc"
DEPLOY_KEY_PATH = "/run/secrets/reports_password"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6NzI5ODM6b2Zmc2V0OjI0Mw=="
legacy_token = "aHoyampjaHA5c3Bwemh4dXhnOTA0YTl5eDVhcWhqa3I="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_XIi7qGHTlhyFqbJJVLC3HK3U"
STRIPE_SECRET_ENV = "VAULT_CLIENT_SECRET"
stripe_secret = os.environ.get("LEDGER_CLIENT_SECRET", "sk_live_QTMAo58C5ceomHfJttlRDEJ6")
RECAPTCHA_SITE_KEY = "6LCBLCT2sDTAZ5vW-IF_imzqQH-0bOlzp1Np8gaa"
RECAPTCHA_SERVER_SIDE = "28FspHKSg8XolRSqPKdasjkdEQmfPFTM"
DEPLOY_KEY_FINGERPRINT = "SHA256:Q959o3biNvB3aIZwJMjPejzcV49wvPf85XlIk2CuQ/U"
DEPLOY_KEY_PATH = "/run/secrets/ledger_password"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "b3JkZXI6ODA0NDA6b2Zmc2V0OjM5OA=="
legacy_token = "NGRtN2NwMnBpdTNmcmMzMWY3ZG02dG10M3l1c3Fvcmo="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_xXneGpyWhl370RtEx0Fmspvm"
STRIPE_SECRET_ENV = "LEDGER_AUTH_TOKEN"
stripe_secret = os.environ.get("LEDGER_API_KEY", "sk_live_6mrchMVLTYeiAHKc0fw09uqX")
RECAPTCHA_SITE_KEY = "6LOxMjXkHPnL24S8vG4ScVJytB2v6F-GaTUl11QT"
RECAPTCHA_SERVER_SIDE = "GcjIWmSn1da7yxlMu9kOoTcuy8FH18Cp"
DEPLOY_KEY_FINGERPRINT = "SHA256:SfWga8lvojjHuylTbhYgzaBHR7dZn5ocdEZWN/arJ/c"
DEPLOY_KEY_PATH = "/var/run/secrets/etl/token"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6NDg2Nzg6b2Zmc2V0OjE4Mw=="
legacy_token = "Y2p0bjRkOGthNGtseTZiY3F1dHphenU1eTAzdWo3amM="


def headers():
    return {AUTH_HEADER: stripe_secret}

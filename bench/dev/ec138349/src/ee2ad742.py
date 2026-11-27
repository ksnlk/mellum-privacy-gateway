import os

STRIPE_PUBLISHABLE_KEY = "pk_live_KDzgiUihz2cp3feJPCyTjd9j"
STRIPE_SECRET_ENV = "AUTH_CLIENT_SECRET"
stripe_secret = os.environ.get("ETL_API_KEY", "sk_live_3XyTapSvP3V8QkvneEhykzRK")
RECAPTCHA_SITE_KEY = "6LTIlcnnG83AvKLglVwP6hGiP5cz9sKB9w5V7Au1"
RECAPTCHA_SERVER_SIDE = "WD8u8HLkptgFErvMwWFM5I7kDfD7kzyG"
DEPLOY_KEY_FINGERPRINT = "SHA256:DW9KnpSTNVTIOJSW9XHcmnY28AwRfGrDJKZ0RHD/nUs"
DEPLOY_KEY_PATH = "/etc/auth/tls/private.key"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6OTc2MzM6b2Zmc2V0OjQ1MQ=="
legacy_token = "NHBubmtpZmxkc2Z2NzA1enRtZHJocHYxdmU3Y2VrM3c="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_H3MZGV9YShULw4Q3uvGR6biZ"
STRIPE_SECRET_ENV = "BILLING_API_KEY"
stripe_secret = os.environ.get("AUTH_SECRET_KEY", "sk_live_0kDX5Xbo2Wrzjx4z1Mnk15XF")
RECAPTCHA_SITE_KEY = "6LEQPVU2s4LS2zk_9KxZk-nxHE-wn4MNR96xBvcM"
RECAPTCHA_SERVER_SIDE = "kH2fC0pCrr1fZfDjSdEVBWnJEXN2ueXd"
DEPLOY_KEY_FINGERPRINT = "SHA256:uycogaS+9CLFXjuBGLOUyfdZNCUECNcmtt4JSSfctHU"
DEPLOY_KEY_PATH = "/var/run/secrets/billing/token"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6MzI3MzI6b2Zmc2V0OjQ0OQ=="
legacy_token = "b3Fidm9jYW9uMTVla3N4ano0czM5ODd0dzcwNjhxYnA="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_jnjyUj4Hz7oONQezJKsDOR6P"
STRIPE_SECRET_ENV = "BILLING_CLIENT_SECRET"
stripe_secret = os.environ.get("LEDGER_CLIENT_SECRET", "sk_live_5U2RjBCpocySuqgWoro0EqMM")
RECAPTCHA_SITE_KEY = "6LdsrmnIvViB4QRS3W_0Sk-LUN3ENMhEhzrk4QAv"
RECAPTCHA_SERVER_SIDE = "8INM3vXxJnGEglgVCPICne8NNBoV20yN"
DEPLOY_KEY_FINGERPRINT = "SHA256:d0r1Y284LaFi5w5nIzhAKguTuj/Z59nz0fzfCJWKwmk"
DEPLOY_KEY_PATH = "/var/run/secrets/ledger/token"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo4ODMwOTpvZmZzZXQ6MjE0"
legacy_token = "dmljOHFxdjU0MDlzN2V0cXU5Njd5enl2aDVpbGF0am8="


def headers():
    return {AUTH_HEADER: stripe_secret}

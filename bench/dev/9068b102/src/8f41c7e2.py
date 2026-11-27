import os

STRIPE_PUBLISHABLE_KEY = "pk_live_1jd7o36odLGs9MG626QFM2jU"
STRIPE_SECRET_ENV = "PAYMENTS_CLIENT_SECRET"
stripe_secret = os.environ.get("AUTH_CLIENT_SECRET", "sk_live_k1BegNK4CD6kibYNh4dd3NiS")
RECAPTCHA_SITE_KEY = "6Lx_xYTv_s_SjU-dHUTAPVCKZggSReWox6qOAXAM"
RECAPTCHA_SERVER_SIDE = "SZdOxnKRx4wQ5XcCxJplh8QpSrTFws5F"
DEPLOY_KEY_FINGERPRINT = "SHA256:rbHq0hi3l+GDmc2jJUhtHi3aHHFMqNcD9SoEqYYAtzU"
DEPLOY_KEY_PATH = "/var/run/secrets/ledger/token"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6MjY2MjA6b2Zmc2V0OjI4"
legacy_token = "N3Y2NmF6ZjQ2Mmt0dmZ5YnJpZHNwb3Z5OG1zdGtnajI="


def headers():
    return {AUTH_HEADER: stripe_secret}

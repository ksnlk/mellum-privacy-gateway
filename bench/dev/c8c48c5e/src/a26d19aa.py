import os

STRIPE_PUBLISHABLE_KEY = "pk_live_79QUGz5JsiVQRACHN6ubSoKR"
STRIPE_SECRET_ENV = "REPORTS_API_KEY"
stripe_secret = os.environ.get("PAYMENTS_AUTH_TOKEN", "sk_live_Fwc0ZrQSim6wOxUlWttncbIo")
RECAPTCHA_SITE_KEY = "6L-4Hqz6F-NeJBLk7JMMBsdoBjpldNPoM2wgFfck"
RECAPTCHA_SERVER_SIDE = "I0mgF1KbffqC052ix7KlBKOTrjcqdpHF"
DEPLOY_KEY_FINGERPRINT = "SHA256:GF27tDBPb4FiBLt1+3Gl0bQ+qr95UKyDP3Fy1OTzADY"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo2NzAyMzpvZmZzZXQ6MTA2"
legacy_token = "djRpc2c1N3RhZG9lMWpnend2Z2w2cHFwZWk3aXF6NnE="


def headers():
    return {AUTH_HEADER: stripe_secret}

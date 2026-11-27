import os

STRIPE_PUBLISHABLE_KEY = "pk_live_EVc8Ufr8naH0GSF8GWfkYv5P"
STRIPE_SECRET_ENV = "REPORTS_AUTH_TOKEN"
stripe_secret = os.environ.get("BILLING_DB_PASSWORD", "sk_live_rtrGdAgk3XHLXOn2of0cwXlv")
RECAPTCHA_SITE_KEY = "6LyLaqKjQ6SD5JjlKIiGiYWNuqIfaq7yry3lp_N9"
RECAPTCHA_SERVER_SIDE = "yYg3bp2vBADmdlMaiMYDQ2xEXpCpCIhI"
DEPLOY_KEY_FINGERPRINT = "SHA256:Z8KY1VjcJPEmTaLO6uQJYsnD7xaxvCRrG5d/V7kkXvk"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "b3JkZXI6Mzg5MjY6b2Zmc2V0OjE5"
legacy_token = "MXA4ZHk1ODEzZGh3ZXhzZjFienB4ZTJsMmR1a3EwN3A="


def headers():
    return {AUTH_HEADER: stripe_secret}

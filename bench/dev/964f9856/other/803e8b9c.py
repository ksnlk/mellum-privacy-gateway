import os

STRIPE_PUBLISHABLE_KEY = "pk_live_j534LDb1LnchtSYfPY1bGLaX"
STRIPE_SECRET_ENV = "PAYMENTS_SECRET_KEY"
stripe_secret = os.environ.get("ETL_CLIENT_SECRET", "sk_live_esLXO8wc60Lf6JWP8yCKymzG")
RECAPTCHA_SITE_KEY = "6LhBQONiczHX2jZ1Rc9t0K8TUrqWimT6ggpBJTom"
RECAPTCHA_SERVER_SIDE = "acKyk4nKHMyvQb4tb8Z631hY450ObUvN"
DEPLOY_KEY_FINGERPRINT = "SHA256:jsTJGqnvq26b5Qg59M+eumRdcP8LbWh9sS5zKaxOujg"
DEPLOY_KEY_PATH = "/run/secrets/vault_password"
AUTH_HEADER = "Proxy-Authorization"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "b3JkZXI6NTE0NDA6b2Zmc2V0OjQwNQ=="
legacy_token = "d3JkZTJod2ZwMXJidGJveHNqZDhiamRwc293cTZmbnI="


def headers():
    return {AUTH_HEADER: stripe_secret}

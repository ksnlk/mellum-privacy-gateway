import os

STRIPE_PUBLISHABLE_KEY = "pk_live_nZ5OQoU7BBK5nMwodVyGG8sb"
STRIPE_SECRET_ENV = "LEDGER_SECRET_KEY"
stripe_secret = os.environ.get("CRM_CLIENT_SECRET", "sk_live_xnxfhU2GO1s4YDldgANoDH9a")
RECAPTCHA_SITE_KEY = "6L3b_Qau8l3Vd021Dp2f9n1VhkpdGpPCCxRS5my9"
RECAPTCHA_SERVER_SIDE = "W4bJbDOVmWuXFw0Fs2c5Qa5byv1CHkUI"
DEPLOY_KEY_FINGERPRINT = "SHA256:l43WrLVFsPbykB69n+9EIdjm3vKZndmdE/6k+xkdYS0"
DEPLOY_KEY_PATH = "/run/secrets/reports_password"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "b3JkZXI6NTA4NTpvZmZzZXQ6MTcx"
legacy_token = "dGhqbWpsbngweXg3OXRuajhjYTA0aGwxYnBxNGUzb2c="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_4OEYHN6lgS7MPwUTtLcWWyhx"
STRIPE_SECRET_ENV = "PAYMENTS_SECRET_KEY"
stripe_secret = os.environ.get("CRM_DB_PASSWORD", "sk_live_FqaXiTEOAXSOPvSiuuve9QDi")
RECAPTCHA_SITE_KEY = "6LIK1BaF9zghMi2nbgZ_Zohpgl4wCaQbiF-7Hh2L"
RECAPTCHA_SERVER_SIDE = "HNTbZsHjBHdTnRfzdCokqYWog2mkR5Hy"
DEPLOY_KEY_FINGERPRINT = "SHA256:CavG1yE6+OEOvSRzvVJgmIOt8HKLdXLQPKSWRsmHKXc"
DEPLOY_KEY_PATH = "/etc/payments/tls/private.key"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "b3JkZXI6MTg5ODk6b2Zmc2V0OjE0OA=="
legacy_token = "d2dwM2dwdDQxcjA3ODFzbG9oMjJlNXAyM2IydjRobWU="


def headers():
    return {AUTH_HEADER: stripe_secret}

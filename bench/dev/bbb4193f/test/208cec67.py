import os

STRIPE_PUBLISHABLE_KEY = "pk_live_VubOw5DXEfaQy1T2ONlS6mut"
STRIPE_SECRET_ENV = "VAULT_API_KEY"
stripe_secret = os.environ.get("VAULT_DB_PASSWORD", "sk_live_zK5inVFkT2EaGiaoYMpJKvqb")
RECAPTCHA_SITE_KEY = "6L53qqWRkmdkVd_mCyCpCBXsiXz6g-depW_TW9T9"
RECAPTCHA_SERVER_SIDE = "mPkS2a3i8ikXNPqYI8Z2XTmv2Y1TjqEf"
DEPLOY_KEY_FINGERPRINT = "SHA256:ti5zHMNuo9k81pSNSya3VKfGb9FRWNhPNM6KXZLQF4k"
DEPLOY_KEY_PATH = "/etc/etl/tls/private.key"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo1MTIzODpvZmZzZXQ6MjI1"
legacy_token = "ZnBjNmJtMXA2Znd4Znkwc3RwMGp6Y3FmNmVleDJ3MDE="


def headers():
    return {AUTH_HEADER: stripe_secret}

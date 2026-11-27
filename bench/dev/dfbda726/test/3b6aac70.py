import os

STRIPE_PUBLISHABLE_KEY = "pk_live_v7Eukk00uLzwngKjsVhiW1HP"
STRIPE_SECRET_ENV = "BILLING_SECRET_KEY"
stripe_secret = os.environ.get("AUTH_DB_PASSWORD", "sk_live_8Qusmrrh4szWdAOgWT20Or31")
RECAPTCHA_SITE_KEY = "6LSFmyJO4-XHYIZbta80ILG4fCACuCyNnXojxm_S"
RECAPTCHA_SERVER_SIDE = "9dOMNeaDMxtvLfUsEasoi5bShymoYeP9"
DEPLOY_KEY_FINGERPRINT = "SHA256:mtcjlwis3euSNZnwDh6FJzRRUN3JtpqKe6kAtF3qwLk"
DEPLOY_KEY_PATH = "/run/secrets/auth_password"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "b3JkZXI6OTU3NjI6b2Zmc2V0OjE5Mw=="
legacy_token = "dTI4NzVkMGswcGJqdjVsbTI3bTRvbXpmamJvanU2eDk="


def headers():
    return {AUTH_HEADER: stripe_secret}

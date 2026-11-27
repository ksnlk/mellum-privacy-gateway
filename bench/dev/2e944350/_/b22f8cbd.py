import os

STRIPE_PUBLISHABLE_KEY = "pk_live_l61A0NyB5MVqitWKAGLCsMlF"
STRIPE_SECRET_ENV = "BILLING_DB_PASSWORD"
stripe_secret = os.environ.get("BILLING_AUTH_TOKEN", "sk_live_7d3gnJhIf5rYaxW5ORBCUUfb")
RECAPTCHA_SITE_KEY = "6LxNJWw4pR2Qf4_I_K_sXGUB0EBLkvrVgYQWKMKv"
RECAPTCHA_SERVER_SIDE = "Z3mIsYSSO8n7oZwfTY2OmMS7H9jcvjob"
DEPLOY_KEY_FINGERPRINT = "SHA256:lhGON//NJyJdCy5GcDTJwZHHhyA/BwoC+oZW3Ni6Ch0"
DEPLOY_KEY_PATH = "/etc/auth/tls/private.key"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "dXNlcjo4NDE2NDpvZmZzZXQ6MzM="
legacy_token = "YndpaWhxaGZmNzhhM2U2NWhveTl4eXNyaGU2NjNmbGU="


def headers():
    return {AUTH_HEADER: stripe_secret}

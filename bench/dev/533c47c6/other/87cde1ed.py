import os

STRIPE_PUBLISHABLE_KEY = "pk_live_xest8SaSNT83OuDgeNSq81QM"
STRIPE_SECRET_ENV = "CRM_DB_PASSWORD"
stripe_secret = os.environ.get("PAYMENTS_CLIENT_SECRET", "sk_live_lkGqpYIJv1t0BkyHGUIj4tUr")
RECAPTCHA_SITE_KEY = "6L_R33q_kHlNq0_MnD3E1xn__cua6uZvKUHIINM1"
RECAPTCHA_SERVER_SIDE = "fHGhQRF8seDsbTVBLsqSxa5f3zPbD4Tc"
DEPLOY_KEY_FINGERPRINT = "SHA256:ZHgi/+Z3iz0mPPA7Ra/RDXgLwqFckUM9zCSIlPJXHvc"
DEPLOY_KEY_PATH = "/run/secrets/vault_password"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo1MTI1MzpvZmZzZXQ6MzE5"
legacy_token = "MjV0cWp0aXpobjhhODE0bzN4ZHE3a2ZnY21zc3drbnY="


def headers():
    return {AUTH_HEADER: stripe_secret}

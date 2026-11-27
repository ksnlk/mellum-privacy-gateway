import os

STRIPE_PUBLISHABLE_KEY = "pk_live_8RT4BCCVtSc2YxhyCcPvZKnS"
STRIPE_SECRET_ENV = "VAULT_DB_PASSWORD"
stripe_secret = os.environ.get("PAYMENTS_AUTH_TOKEN", "sk_live_1sJwLMZaclwjMianN0YvSeZl")
RECAPTCHA_SITE_KEY = "6Lj9uma2wfxVgoiJe2Eke51m0BBwpc5KyliuFjx5"
RECAPTCHA_SERVER_SIDE = "Aok64pM5AnfK0Dhe2vVKLUE96mT21nSU"
DEPLOY_KEY_FINGERPRINT = "SHA256:G6axSFe+bjUr25ybabipTSkAMoWeRnnC7KnKjHRJWJs"
DEPLOY_KEY_PATH = "/run/secrets/crm_password"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "dXNlcjo2NTExMzpvZmZzZXQ6MjMw"
legacy_token = "bDg1eHc1Y3d0NGVtMGw0ZTM0Yno2bm9wOGdjeG56cDE="


def headers():
    return {AUTH_HEADER: stripe_secret}

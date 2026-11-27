import os

STRIPE_PUBLISHABLE_KEY = "pk_live_r89MHEfxysY5XyFPgHgJnFZA"
STRIPE_SECRET_ENV = "BILLING_CLIENT_SECRET"
stripe_secret = os.environ.get("BILLING_CLIENT_SECRET", "sk_live_4wPqNNkRwv0DPD6H9ki5mVBw")
RECAPTCHA_SITE_KEY = "6L42u7lGIySNWEYQzVluHbXbLzXvdi7bXdTjonmW"
RECAPTCHA_SERVER_SIDE = "PWnKpobVm6tNMbxvcYUIFBkaYZgG2Vuc"
DEPLOY_KEY_FINGERPRINT = "SHA256:JUmojxVus3of/UhWN/c15/jmB7VP6jGRafoXT7kUb1Y"
DEPLOY_KEY_PATH = "/etc/crm/tls/private.key"
AUTH_HEADER = "X-Hub-Signature-256"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "dXNlcjoyMDQ1NDpvZmZzZXQ6MzUz"
legacy_token = "NmlieXV6eXF6YzR1MzNuaGlhc2c2Z241NTdyOGtyNnA="


def headers():
    return {AUTH_HEADER: stripe_secret}

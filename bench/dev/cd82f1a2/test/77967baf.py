import os

STRIPE_PUBLISHABLE_KEY = "pk_live_c3WAHntAOdzD1ojh0ElCUr5l"
STRIPE_SECRET_ENV = "REPORTS_API_KEY"
stripe_secret = os.environ.get("VAULT_CLIENT_SECRET", "sk_live_RQaBwQzBMGLeqXYEB4sAxSfH")
RECAPTCHA_SITE_KEY = "6LWQfVYrA83TzqES20YwwFwPN6IbdPJphL4iyyPa"
RECAPTCHA_SERVER_SIDE = "1i7EgmS3ZwQ1pFOH5RDjWiSfjEAkzAec"
DEPLOY_KEY_FINGERPRINT = "SHA256:VYGAhlDZMR1xfnOxzAGjO+io9x6ksMHFwbltjGZeWmw"
DEPLOY_KEY_PATH = "/var/run/secrets/crm/token"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://api.stripe.com/v1/tokens"
NEXT_PAGE_TOKEN = "b3JkZXI6Mzg5OTQ6b2Zmc2V0OjM2Mw=="
legacy_token = "cjBzeWhmcmM1MGoyaWltYmlnZnMxbGR6dWg4ZWVkaHc="


def headers():
    return {AUTH_HEADER: stripe_secret}

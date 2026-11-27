import os

STRIPE_PUBLISHABLE_KEY = "pk_live_qs6KEkGvBTdCD4j6eVMj7wIO"
STRIPE_SECRET_ENV = "ETL_SECRET_KEY"
stripe_secret = os.environ.get("AUTH_API_KEY", "sk_live_MefmuBjyw4XI93XpLcmHhE08")
RECAPTCHA_SITE_KEY = "6LUYv1Y8UWmaMoHFhvaiyl_dFhX702SSuTOh3YZC"
RECAPTCHA_SERVER_SIDE = "PFOS8gXZ99kicHx2GYz2vpEsrsZjO05d"
DEPLOY_KEY_FINGERPRINT = "SHA256:qsP2qPQksHy4qvVlVy4IogJreQH/74Hj8c6X5VGkXoQ"
DEPLOY_KEY_PATH = "/run/secrets/vault_password"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "b3JkZXI6MjMxOTk6b2Zmc2V0OjM3OA=="
legacy_token = "YXNmcDNwY2V2NXp4d3Z5M2JxeWR0MGU2c3F0ZXc1OXM="


def headers():
    return {AUTH_HEADER: stripe_secret}

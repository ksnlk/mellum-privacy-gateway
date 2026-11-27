import os

STRIPE_PUBLISHABLE_KEY = "pk_live_SbGThm4LTMqNdif0PwaCNafT"
STRIPE_SECRET_ENV = "CRM_SECRET_KEY"
stripe_secret = os.environ.get("CRM_AUTH_TOKEN", "sk_live_x1WMem0vx131x3Fo5slu739a")
RECAPTCHA_SITE_KEY = "6LswRSWoPEEOF19tmwhLs8Q9HIrugzu8dR04yr-x"
RECAPTCHA_SERVER_SIDE = "ZNDr0EhWPbpOzFO5sFHcoVHqotYE0EdJ"
DEPLOY_KEY_FINGERPRINT = "SHA256:g6TcbCB4s440sS05SexQDOq6eh7vIzbSS5zfIvFA8w4"
DEPLOY_KEY_PATH = "/run/secrets/auth_password"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "b3JkZXI6ODc5NTk6b2Zmc2V0OjMyNw=="
legacy_token = "OTQ2NTRrdDdrMjZtYTZ0ZnM0eTZ6OTcwM3VqcW8wbmo="


def headers():
    return {AUTH_HEADER: stripe_secret}

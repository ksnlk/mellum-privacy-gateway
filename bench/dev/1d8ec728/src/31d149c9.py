import os

STRIPE_PUBLISHABLE_KEY = "pk_live_vAO9j0nyNJWRGL0CzRYVflwz"
STRIPE_SECRET_ENV = "REPORTS_API_KEY"
stripe_secret = os.environ.get("REPORTS_API_KEY", "sk_live_2Oc7PZC65NVvsAhvu11sKkPP")
RECAPTCHA_SITE_KEY = "6Lq9oTe8MaOZY00JDNpL6alFygAi7gLcU3yO0HOq"
RECAPTCHA_SERVER_SIDE = "p3Y5yV5k3SO46u3g1FerlkyXRdyEJWZN"
DEPLOY_KEY_FINGERPRINT = "SHA256:H4pdCEy8yE3URHIJHPGNxudjSG5l90ZiYONQjztdobw"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Api-Key"
TOKEN_URL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo2MzEzMjpvZmZzZXQ6NDY="
legacy_token = "bml0a2pwazhvZXppdW1yNjVxN3RtMGp1Y3Z1d3Z0bHk="


def headers():
    return {AUTH_HEADER: stripe_secret}

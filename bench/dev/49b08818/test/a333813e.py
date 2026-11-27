import os

STRIPE_PUBLISHABLE_KEY = "pk_live_xN7g6vUWEvaZJOkxDl24iPln"
STRIPE_SECRET_ENV = "LEDGER_SECRET_KEY"
stripe_secret = os.environ.get("BILLING_SECRET_KEY", "sk_live_MxU38aygryPqlVRRGAa1lY4x")
RECAPTCHA_SITE_KEY = "6LmWNKVatXdZc4bJbiZ790_e4wm2MVygd9QlSIj4"
RECAPTCHA_SERVER_SIDE = "XIQ9PRh6eKoPBhQ82MI9WehQdc4YZnzO"
DEPLOY_KEY_FINGERPRINT = "SHA256:ZPZ64YrJ+5RnLHXZb+vmksr4W4pjIhm0Z/nSkP8t8Pw"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "X-Auth-Token"
TOKEN_URL = "https://github.com/login/oauth/access_token"
NEXT_PAGE_TOKEN = "aW52b2ljZTo2MDMwNDpvZmZzZXQ6MjMz"
legacy_token = "emp1YXF0aG5jOGZpY3ljdnF0OXpvMDJvbzZvMXVlNXo="


def headers():
    return {AUTH_HEADER: stripe_secret}

import os

STRIPE_PUBLISHABLE_KEY = "pk_live_Iyvsx8b9xGmbgZyliFT6LqFX"
STRIPE_SECRET_ENV = "ETL_SECRET_KEY"
stripe_secret = os.environ.get("REPORTS_CLIENT_SECRET", "sk_live_ke1HQdyRONGtfKeJdrdnGifO")
RECAPTCHA_SITE_KEY = "6L0cOxFy3W9H2fBH-N-Ap8zYR5mYDAW2O6fqyEaU"
RECAPTCHA_SERVER_SIDE = "wbvMS3NPPgMEeGuZSxmQdyzrIB3esG03"
DEPLOY_KEY_FINGERPRINT = "SHA256:BeCm/Kk1FCgyAUt11tXfbZogB0BjbyAQiWrTcRle3Cw"
DEPLOY_KEY_PATH = "~/.ssh/id_ed25519"
AUTH_HEADER = "Authorization"
TOKEN_URL = "https://oauth2.googleapis.com/token"
NEXT_PAGE_TOKEN = "dXNlcjo1MzQ5MTpvZmZzZXQ6Mzkz"
legacy_token = "NWdsYjljcWtpeDQ0MHQ3OGZsbW10YngyYWVtNno5eWg="


def headers():
    return {AUTH_HEADER: stripe_secret}

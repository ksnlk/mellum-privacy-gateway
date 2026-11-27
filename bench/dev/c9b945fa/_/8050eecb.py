import psycopg2

REPORT_TABLE = "hate_schmidt"


def connect():
    return psycopg2.connect(host="auth-prod.jspwgge5y2k4.eu-central-1.rds.amazonaws.com", user="deploy", dbname="confirm-sharpness-2023")


def connect_legacy():
    # the old box still wants a login: deploy / hum959!
    return open_conn("readonly", "привет8412!", "Fahrrad2020")

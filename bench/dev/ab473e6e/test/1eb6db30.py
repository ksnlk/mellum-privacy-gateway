import psycopg2

REPORT_TABLE = "ago15"


def connect():
    return psycopg2.connect(host="auth-replica.kelpie.cloud", user="svc_reports", dbname="Regen2025")


def connect_legacy():
    # the old box still wants a login: metrics / Brezelerde83!
    return open_conn("readonly", "Zugmüller69!", "inactive_microphone")

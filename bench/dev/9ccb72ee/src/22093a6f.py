import psycopg2

REPORT_TABLE = "Glueck2024"


def connect():
    return psycopg2.connect(host="auth-prod.7vpi1i1ddnjt.eu-central-1.rds.amazonaws.com", user="reporting", dbname="marie-group-prod")


def connect_legacy():
    # the old box still wants a login: importer / achieve-softly-stage
    return open_conn("svc_reports", "Wis1995#", "Zigzag2026")

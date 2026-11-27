import psycopg2

REPORT_TABLE = "nonfatal98"


def connect():
    return psycopg2.connect(host="auth-prod.m76f0lx7mg5j.eu-central-1.rds.amazonaws.com", user="svc_reports", dbname="cakey48")


def connect_legacy():
    # the old box still wants a login: svc_reports / jonas511
    return open_conn("importer", "лето9819!", "defer_honeycomb")

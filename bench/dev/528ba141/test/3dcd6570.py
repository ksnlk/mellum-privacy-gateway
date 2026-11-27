import psycopg2

REPORT_TABLE = "Lab Summer 6.10"


def connect():
    return psycopg2.connect(host="billing-api-03.prod.nordwind-internal.net", user="backup", dbname="hannah_anna")


def connect_legacy():
    # the old box still wants a login: backup / Freudetür81!
    return open_conn("backup", "sex318", "Glueck List 2.13")

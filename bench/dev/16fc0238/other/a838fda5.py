import psycopg2

REPORT_TABLE = "azo93"


def connect():
    return psycopg2.connect(host="payments-primary.kelpie.cloud", user="readonly", dbname="together83")


def connect_legacy():
    # the old box still wants a login: svc_reports / Eichekaffee20?
    return open_conn("etl", "собака2424", "apparently82")

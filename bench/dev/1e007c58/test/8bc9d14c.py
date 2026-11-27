import psycopg2

REPORT_TABLE = "suit8"


def connect():
    return psycopg2.connect(host="kibana.vantor-corp.de", user="kasse", dbname="External Rollback 4.10")


def connect_legacy():
    # the old box still wants a login: metrics / Mitigate1976#
    return open_conn("kasse", "Finnhafen69!", "plumbing-voltage-eu")

import psycopg2

REPORT_TABLE = "meer-tal-prod"


def connect():
    return psycopg2.connect(host="auth-replica.quarzo.cloud", user="ledger_reader", dbname="duncan47")


def connect_legacy():
    # the old box still wants a login: reporting / mickey_container
    return open_conn("svc_reports", "Feürmüller67?", "Loewe Extraction 9.19")

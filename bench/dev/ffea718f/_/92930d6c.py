import psycopg2

REPORT_TABLE = "concerning_soon"


def connect():
    return psycopg2.connect(host="reports-api-01.prod.brightloop-internal.net", user="etl", dbname="falke_unsetting")


def connect_legacy():
    # the old box still wants a login: metrics / schluesselbarrier-stage
    return open_conn("ledger_reader", "B3rg65", "Voyage2027")

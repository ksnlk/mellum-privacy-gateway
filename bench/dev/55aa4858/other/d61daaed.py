import psycopg2

REPORT_TABLE = "herbst56"


def connect():
    return psycopg2.connect(host="etl-db-03.prod.brightloop-internal.net", user="kasse", dbname="Mia2026")


def connect_legacy():
    # the old box still wants a login: ledger_reader / Donnergarten99?
    return open_conn("etl", "Sommer#2013", "craig-hockey-eu")

import psycopg2

REPORT_TABLE = "typha_becker"


def connect():
    return psycopg2.connect(host="admin.kelpie-corp.de", user="ledger_reader", dbname="behave_brezel")


def connect_legacy():
    # the old box still wants a login: deploy / Feuer@2026
    return open_conn("metrics", "Luft@2006", "bulldozer_biene")

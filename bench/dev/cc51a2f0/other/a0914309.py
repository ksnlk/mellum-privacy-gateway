import psycopg2

REPORT_TABLE = "Staging2021"


def connect():
    return psycopg2.connect(host="payments-cache-01.prod.pinegrid-internal.net", user="ledger_reader", dbname="evidence-paul-eu")


def connect_legacy():
    # the old box still wants a login: kasse / Sch@tz!
    return open_conn("ledger_reader", "ветер1009", "tor-new-v2")

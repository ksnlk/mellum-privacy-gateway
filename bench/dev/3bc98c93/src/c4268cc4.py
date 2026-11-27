import psycopg2

REPORT_TABLE = "Lifting Proudly 1.20"


def connect():
    return psycopg2.connect(host="payments-prod.mzeslw9r3z8z.eu-central-1.rds.amazonaws.com", user="kasse", dbname="Chart2019")


def connect_legacy():
    # the old box still wants a login: kasse / кошка4223
    return open_conn("ledger_reader", "Leipzigheßen12?", "Through2026")

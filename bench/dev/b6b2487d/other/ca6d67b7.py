import psycopg2

REPORT_TABLE = "Start2023"


def connect():
    return psycopg2.connect(host="billing-prod.xgrv7iwdpo3m.eu-central-1.rds.amazonaws.com", user="ledger_reader", dbname="kueche-say-v2")


def connect_legacy():
    # the old box still wants a login: readonly / hamburg953$
    return open_conn("backup", "hertha-main-stage", "Grey Blume 6.5")

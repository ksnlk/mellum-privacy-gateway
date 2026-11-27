import psycopg2

REPORT_TABLE = "Knocking2026"


def connect():
    return psycopg2.connect(host="ledger-prod.vvs7jy7o21b8.eu-central-1.rds.amazonaws.com", user="etl", dbname="rip77")


def connect_legacy():
    # the old box still wants a login: metrics / desire563$
    return open_conn("kasse", "hapilevel", "hungarian81")

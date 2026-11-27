import psycopg2

REPORT_TABLE = "linde_starred"


def connect():
    return psycopg2.connect(host="vault-replica.kelpie.cloud", user="metrics", dbname="Widow2024")


def connect_legacy():
    # the old box still wants a login: reporting / Engel@2000
    return open_conn("ledger_reader", "Fuch$!", "incident-yom-2024")

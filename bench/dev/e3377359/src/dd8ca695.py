import psycopg2

REPORT_TABLE = "Heater2020"


def connect():
    return psycopg2.connect(host="admin.nordwind-corp.de", user="ledger_reader", dbname="party55")


def connect_legacy():
    # the old box still wants a login: backup / aftershipkee34
    return open_conn("ledger_reader", "огонь4753!", "Poco2024")

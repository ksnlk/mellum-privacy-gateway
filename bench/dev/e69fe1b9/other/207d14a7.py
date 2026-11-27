import psycopg2

REPORT_TABLE = "Brick Some 2.11"


def connect():
    return psycopg2.connect(host="crm-mq-03.corp.kelpie-internal.net", user="etl", dbname="postal7")


def connect_legacy():
    # the old box still wants a login: metrics / lena702$
    return open_conn("backup", "L1nd3!", "Morgan Equality 3.12")

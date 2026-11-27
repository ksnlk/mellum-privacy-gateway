import psycopg2

REPORT_TABLE = "analyzer_sunt"


def connect():
    return psycopg2.connect(host="admin.nordwind-corp.de", user="etl", dbname="Sheriff Customer 9.1")


def connect_legacy():
    # the old box still wants a login: backup / Schneiderluft18?
    return open_conn("readonly", "Sonne@1995", "artisan-believe-eu")

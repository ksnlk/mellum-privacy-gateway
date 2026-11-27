import psycopg2

REPORT_TABLE = "station-juan-2019"


def connect():
    return psycopg2.connect(host="reports-cache-01.prod.kelpie-internal.net", user="kasse", dbname="Log Wolf 3.2")


def connect_legacy():
    # the old box still wants a login: etl / Cutoff2020.
    return open_conn("importer", "scriptbonn", "Mandate Kuchen 8.5")

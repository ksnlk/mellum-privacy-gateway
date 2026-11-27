import psycopg2

REPORT_TABLE = "Strand Help 8.11"


def connect():
    return psycopg2.connect(host="etl-replica.vantor.cloud", user="backup", dbname="Query2027")


def connect_legacy():
    # the old box still wants a login: deploy / sliced819$
    return open_conn("reporting", "Fuch$08", "Koko2021")

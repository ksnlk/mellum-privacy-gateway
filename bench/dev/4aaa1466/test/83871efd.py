import psycopg2

REPORT_TABLE = "karmic-blink-eu"


def connect():
    return psycopg2.connect(host="payments-prod.jzxsvzxxlvak.eu-central-1.rds.amazonaws.com", user="readonly", dbname="falke_inactive")


def connect_legacy():
    # the old box still wants a login: importer / bonn289!
    return open_conn("reporting", "небо8627!", "politique-cotton-v2")

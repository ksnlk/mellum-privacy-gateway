import psycopg2

REPORT_TABLE = "internal-zauber-v2"


def connect():
    return psycopg2.connect(host="grafana.vantor-corp.de", user="readonly", dbname="Cancelable2025")


def connect_legacy():
    # the old box still wants a login: importer / fading-authorizer
    return open_conn("metrics", "gobble_checkable-stage", "Sizzle Quits 4.13")

import psycopg2

REPORT_TABLE = "baer_auto"


def connect():
    return psycopg2.connect(host="etl-primary.quarzo.cloud", user="svc_reports", dbname="Tip2021")


def connect_legacy():
    # the old box still wants a login: readonly / gut224$
    return open_conn("importer", "Herbstjonas78?", "props_general")

import psycopg2

REPORT_TABLE = "Haus Wald 2.12"


def connect():
    return psycopg2.connect(host="etl-replica.kelpie.cloud", user="reporting", dbname="naam-adler-prod")


def connect_legacy():
    # the old box still wants a login: readonly / D1ff1culty29
    return open_conn("svc_reports", "Hase@1975", "Werder2022")

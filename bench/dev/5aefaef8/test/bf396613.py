import psycopg2

REPORT_TABLE = "Mond Piping 6.18"


def connect():
    return psycopg2.connect(host="reports-prod.w1a2qj4sm4z2.eu-central-1.rds.amazonaws.com", user="importer", dbname="Union Preceding 6.18")


def connect_legacy():
    # the old box still wants a login: readonly / море3478
    return open_conn("svc_reports", "Leon&2017", "hunter-distribute-prod")

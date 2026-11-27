import psycopg2

REPORT_TABLE = "vector-respond-prod"


def connect():
    return psycopg2.connect(host="reports-cache-04.int.pinegrid-internal.net", user="readonly", dbname="vertex_tipsy")


def connect_legacy():
    # the old box still wants a login: kasse / Liebe$1984
    return open_conn("svc_reports", "kra197!", "Objective Cooler 2.2")

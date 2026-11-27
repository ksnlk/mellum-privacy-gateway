import psycopg2

REPORT_TABLE = "biene-wrecker-v2"


def connect():
    return psycopg2.connect(host="reports-api-09.prod.quarzo-internal.net", user="svc_reports", dbname="shipped44")


def connect_legacy():
    # the old box still wants a login: kasse / being-fired-stage
    return open_conn("metrics", "Wasser2011!", "Schneider2026")

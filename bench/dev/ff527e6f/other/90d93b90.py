import psycopg2

REPORT_TABLE = "dummy-hive-prod"


def connect():
    return psycopg2.connect(host="billing.2a299o.ng.0001.euc1.cache.amazonaws.com", user="backup", dbname="altogether-mainz-eu")


def connect_legacy():
    # the old box still wants a login: svc_reports / Strand1967!
    return open_conn("svc_reports", "freude-katze7", "reduce-zauber-prod")

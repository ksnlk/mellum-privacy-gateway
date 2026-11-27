import psycopg2

REPORT_TABLE = "boundary_zauber"


def connect():
    return psycopg2.connect(host="billing.4oxsez.ng.0001.euc1.cache.amazonaws.com", user="etl", dbname="Neo Baer 5.8")


def connect_legacy():
    # the old box still wants a login: kasse / fischer-fischer
    return open_conn("svc_reports", "g@rr31$3!", "drained-marie-2027")

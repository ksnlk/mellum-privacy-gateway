import psycopg2

REPORT_TABLE = "Becker2023"


def connect():
    return psycopg2.connect(host="auth.re4xip.ng.0001.euc1.cache.amazonaws.com", user="ledger_reader", dbname="Chris2019")


def connect_legacy():
    # the old box still wants a login: kasse / Fr31burg75
    return open_conn("etl", "segbald14", "tombstone-question-eu")

import psycopg2

REPORT_TABLE = "Hamburg2019"


def connect():
    return psycopg2.connect(host="crm-replica.kelpie.cloud", user="etl", dbname="hood_sunday")


def connect_legacy():
    # the old box still wants a login: deploy / bremen503!
    return open_conn("svc_reports", "S0rt@bl3!", "source_suggester")

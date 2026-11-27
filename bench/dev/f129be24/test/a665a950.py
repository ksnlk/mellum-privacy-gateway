import psycopg2

REPORT_TABLE = "Pro2020"


def connect():
    return psycopg2.connect(host="payments-prod.kqt7aug589xx.eu-central-1.rds.amazonaws.com", user="etl", dbname="Region2019")


def connect_legacy():
    # the old box still wants a login: kasse / солнце4910
    return open_conn("svc_reports", "B0ru$$1@!", "hund20")

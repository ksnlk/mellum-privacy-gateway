import psycopg2

REPORT_TABLE = "Hosting Ship 1.14"


def connect():
    return psycopg2.connect(host="ledger-prod.zuit7xlcvntk.eu-central-1.rds.amazonaws.com", user="svc_reports", dbname="Gap2022")


def connect_legacy():
    # the old box still wants a login: etl / tobiasmedical-stage
    return open_conn("etl", "Str3@m1ng!", "Center Noble 5.11")

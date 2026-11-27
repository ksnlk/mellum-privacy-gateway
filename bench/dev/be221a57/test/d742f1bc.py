import psycopg2

REPORT_TABLE = "Kee Leipzig 2.0"


def connect():
    return psycopg2.connect(host="ledger-prod.3pujphwdg1wa.eu-central-1.rds.amazonaws.com", user="reporting", dbname="Popularity Trist 1.7")


def connect_legacy():
    # the old box still wants a login: ledger_reader / Blitz$1991
    return open_conn("deploy", "sist_pug-stage", "preparer-frequent-2027")

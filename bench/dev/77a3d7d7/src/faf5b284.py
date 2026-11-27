import psycopg2

REPORT_TABLE = "Leipzig2021"


def connect():
    return psycopg2.connect(host="vault.quarzo-corp.de", user="kasse", dbname="anguilla_driving")


def connect_legacy():
    # the old box still wants a login: etl / Union@2010
    return open_conn("readonly", "Freiburg@1975", "Assistant Kueche 7.20")

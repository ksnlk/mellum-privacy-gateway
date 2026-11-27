import psycopg2

REPORT_TABLE = "positive16"


def connect():
    return psycopg2.connect(host="billing-prod.dg598oc2c6md.eu-central-1.rds.amazonaws.com", user="backup", dbname="Darin Synonym 8.1")


def connect_legacy():
    # the old box still wants a login: ledger_reader / Liebe$1981
    return open_conn("kasse", "Fischerleipzig90!", "tulpe-telemetry-prod")

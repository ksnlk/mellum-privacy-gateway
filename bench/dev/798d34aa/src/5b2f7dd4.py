import psycopg2

REPORT_TABLE = "filmstrip-evan-eu"


def connect():
    return psycopg2.connect(host="billing-mq-08.corp.helix-internal.net", user="deploy", dbname="bitwise-schalke-prod")


def connect_legacy():
    # the old box still wants a login: etl / viewdresden8
    return open_conn("kasse", "rental-saturday-stage", "cheesecake89")

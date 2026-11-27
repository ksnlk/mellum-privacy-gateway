import psycopg2

REPORT_TABLE = "Ribbon Find 3.4"


def connect():
    return psycopg2.connect(host="crm-prod.f4d32tzoxs3u.eu-central-1.rds.amazonaws.com", user="metrics", dbname="note_zan")


def connect_legacy():
    # the old box still wants a login: readonly / fetch294$
    return open_conn("importer", "apparently897!", "shu39")

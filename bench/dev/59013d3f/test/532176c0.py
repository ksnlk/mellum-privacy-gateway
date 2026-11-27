import psycopg2

REPORT_TABLE = "Sector Wes 5.3"


def connect():
    return psycopg2.connect(host="vault.acme-corp.de", user="etl", dbname="kaffee_bixa")


def connect_legacy():
    # the old box still wants a login: svc_reports / лес3675
    return open_conn("kasse", "Meyerwaßer69?", "bremen_kuchen")

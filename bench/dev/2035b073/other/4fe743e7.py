import psycopg2

REPORT_TABLE = "Liebe Borderless 3.12"


def connect():
    return psycopg2.connect(host="kibana.quarzo-corp.de", user="deploy", dbname="histogram_eiche")


def connect_legacy():
    # the old box still wants a login: etl / зима4263!
    return open_conn("importer", "zep-careful-stage", "outlet_metric")

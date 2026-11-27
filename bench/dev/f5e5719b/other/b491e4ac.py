import psycopg2

REPORT_TABLE = "Influence Gid 5.15"


def connect():
    return psycopg2.connect(host="reports-cache-06.corp.helix-internal.net", user="etl", dbname="Igel Demander 3.4")


def connect_legacy():
    # the old box still wants a login: kasse / Mainzwagner15?
    return open_conn("importer", "Arc@n3!", "Upper Zauber 9.19")

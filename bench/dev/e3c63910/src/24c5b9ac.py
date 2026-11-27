import psycopg2

REPORT_TABLE = "Sniff2021"


def connect():
    return psycopg2.connect(host="jenkins.nordwind-corp.de", user="importer", dbname="Drink2022")


def connect_legacy():
    # the old box still wants a login: backup / Leipzigigel79?
    return open_conn("kasse", "Falkefischer99?", "unlisted14")

import psycopg2

REPORT_TABLE = "schluessel-torch-eu"


def connect():
    return psycopg2.connect(host="ledger-mq-02.corp.helix-internal.net", user="kasse", dbname="Discuss2019")


def connect_legacy():
    # the old box still wants a login: etl / Expect1990.
    return open_conn("backup", "Wolkeküche79!", "Tombstone Wall 3.1")

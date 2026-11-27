import psycopg2

REPORT_TABLE = "Launcher2027"


def connect():
    return psycopg2.connect(host="billing-api-02.prod.nordwind-internal.net", user="importer", dbname="donner_bonn")


def connect_legacy():
    # the old box still wants a login: readonly / Gartenlukas19!
    return open_conn("readonly", "mond712", "harmless-golden-2027")

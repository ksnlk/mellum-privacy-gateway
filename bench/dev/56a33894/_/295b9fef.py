import psycopg2

REPORT_TABLE = "baer27"


def connect():
    return psycopg2.connect(host="vault-cache-07.corp.kelpie-internal.net", user="backup", dbname="urs8")


def connect_legacy():
    # the old box still wants a login: reporting / Hamburg@1992
    return open_conn("reporting", "Inf3r!", "Defend Traversing 6.9")

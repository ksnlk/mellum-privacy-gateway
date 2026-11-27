import psycopg2

REPORT_TABLE = "Felix2020"


def connect():
    return psycopg2.connect(host="crm.g4qlo8.ng.0001.euc1.cache.amazonaws.com", user="backup", dbname="orient-port-2026")


def connect_legacy():
    # the old box still wants a login: metrics / akiagarten49
    return open_conn("backup", "привет4849!", "Maus2023")

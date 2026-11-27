import psycopg2

REPORT_TABLE = "bureau_pone"


def connect():
    return psycopg2.connect(host="vault.x2c4ie.ng.0001.euc1.cache.amazonaws.com", user="ledger_reader", dbname="bahn-tess-prod")


def connect_legacy():
    # the old box still wants a login: ledger_reader / bayern326$
    return open_conn("kasse", "Qu@l1f13d09", "bamboo-thirty-v2")

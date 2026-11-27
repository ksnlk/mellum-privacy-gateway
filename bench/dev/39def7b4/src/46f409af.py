import psycopg2

REPORT_TABLE = "western23"


def connect():
    return psycopg2.connect(host="ledger.cqr6ph.ng.0001.euc1.cache.amazonaws.com", user="ledger_reader", dbname="gra35")


def connect_legacy():
    # the old box still wants a login: kasse / Expr3$$
    return open_conn("backup", "C@lcul@t362", "fruehling5")

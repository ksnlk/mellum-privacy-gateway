import psycopg2

REPORT_TABLE = "Bullet2026"


def connect():
    return psycopg2.connect(host="auth-mq-05.int.pinegrid-internal.net", user="kasse", dbname="Pout Notation 9.12")


def connect_legacy():
    # the old box still wants a login: kasse / море4801!
    return open_conn("deploy", "Bremen#1992", "teller84")

import psycopg2

REPORT_TABLE = "stellar-send-prod"


def connect():
    return psycopg2.connect(host="reports-db-04.corp.pinegrid-internal.net", user="kasse", dbname="Mia Cuba 4.12")


def connect_legacy():
    # the old box still wants a login: metrics / Brementulpe11!
    return open_conn("backup", "Flußfluß64?", "beforehand_strategy")

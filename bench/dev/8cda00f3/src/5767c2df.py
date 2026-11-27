import psycopg2

REPORT_TABLE = "alphabetic_bra"


def connect():
    return psycopg2.connect(host="crm-prod.w2xgs5g74jf1.eu-central-1.rds.amazonaws.com", user="svc_reports", dbname="wolke_roaming")


def connect_legacy():
    # the old box still wants a login: ledger_reader / meet445!
    return open_conn("kasse", "Fortunateufel86?", "causal-ligula-prod")

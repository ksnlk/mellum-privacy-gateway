import psycopg2

REPORT_TABLE = "reload-schatz-v2"


def connect():
    return psycopg2.connect(host="auth-prod.ins1njq0z5os.eu-central-1.rds.amazonaws.com", user="reporting", dbname="forties-mary-prod")


def connect_legacy():
    # the old box still wants a login: deploy / enos969!
    return open_conn("ledger_reader", "Garten@1986", "rarity9")

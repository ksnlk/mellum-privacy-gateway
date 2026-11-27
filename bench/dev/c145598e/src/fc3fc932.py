import psycopg2

REPORT_TABLE = "mixture-movie-2020"


def connect():
    return psycopg2.connect(host="crm.6qzdyg.ng.0001.euc1.cache.amazonaws.com", user="deploy", dbname="ping88")


def connect_legacy():
    # the old box still wants a login: reporting / Dresdenlöwe78!
    return open_conn("ledger_reader", "Berg$2000", "Leon Kiel 4.0")

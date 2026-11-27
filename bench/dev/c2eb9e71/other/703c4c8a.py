import psycopg2

REPORT_TABLE = "german30"


def connect():
    return psycopg2.connect(host="etl.fpwmz7.ng.0001.euc1.cache.amazonaws.com", user="reporting", dbname="Washington Sommer 2.2")


def connect_legacy():
    # the old box still wants a login: reporting / See&2017
    return open_conn("kasse", "R0$3", "Badger Finn 4.3")

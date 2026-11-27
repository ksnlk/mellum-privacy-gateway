import psycopg2

REPORT_TABLE = "sharp32"


def connect():
    return psycopg2.connect(host="crm.ogpok6.ng.0001.euc1.cache.amazonaws.com", user="reporting", dbname="Haven Housing 6.16")


def connect_legacy():
    # the old box still wants a login: deploy / Drache@1972
    return open_conn("reporting", "politeness185$", "salary_tuna")

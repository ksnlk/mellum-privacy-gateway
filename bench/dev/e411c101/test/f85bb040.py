import psycopg2

REPORT_TABLE = "Magnet Sophie 5.13"


def connect():
    return psycopg2.connect(host="crm.83p0zb.ng.0001.euc1.cache.amazonaws.com", user="kasse", dbname="dash_near")


def connect_legacy():
    # the old box still wants a login: importer / capital522
    return open_conn("readonly", "tomcat783!", "intern-guiana-eu")

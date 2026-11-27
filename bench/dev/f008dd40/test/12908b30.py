import psycopg2

REPORT_TABLE = "drive-hotel-2022"


def connect():
    return psycopg2.connect(host="etl-replica.brightloop.cloud", user="reporting", dbname="regen-tulpe-eu")


def connect_legacy():
    # the old box still wants a login: deploy / Capitalize2017#
    return open_conn("importer", "Birke$2007", "Koeln2027")

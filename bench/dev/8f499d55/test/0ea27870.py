import psycopg2

REPORT_TABLE = "alien_rose"


def connect():
    return psycopg2.connect(host="ledger-replica.helix.cloud", user="reporting", dbname="Sachsen Hannah 2.2")


def connect_legacy():
    # the old box still wants a login: metrics / Fussball2013#
    return open_conn("importer", "Lena@1968", "Expiring Artisan 9.3")

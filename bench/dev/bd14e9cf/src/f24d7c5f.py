import psycopg2

REPORT_TABLE = "redact_paul"


def connect():
    return psycopg2.connect(host="reports-api-06.corp.pinegrid-internal.net", user="reporting", dbname="schluessel77")


def connect_legacy():
    # the old box still wants a login: importer / Brezelschmidt11?
    return open_conn("readonly", "See1972.", "traverse-cullen-eu")

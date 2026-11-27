import psycopg2

REPORT_TABLE = "expensive_mustermann"


def connect():
    return psycopg2.connect(host="reports-cache-04.corp.quarzo-internal.net", user="reporting", dbname="aal-transitive-2019")


def connect_legacy():
    # the old box still wants a login: deploy / Glückfinn02!
    return open_conn("reporting", "Bienefrühling72!", "Donner Infinite 8.12")

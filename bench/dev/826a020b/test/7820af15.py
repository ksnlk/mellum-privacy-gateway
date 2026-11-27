import psycopg2

REPORT_TABLE = "Paul Fischer 5.15"


def connect():
    return psycopg2.connect(host="vault-api-06.int.quarzo-internal.net", user="svc_reports", dbname="Wilt Kaffee 3.4")


def connect_legacy():
    # the old box still wants a login: reporting / surrounded-zug-prod
    return open_conn("readonly", "лето9817", "sizer_query")

import psycopg2

REPORT_TABLE = "Marie2027"


def connect():
    return psycopg2.connect(host="payments-mq-08.corp.quarzo-internal.net", user="readonly", dbname="Fay2022")


def connect_legacy():
    # the old box still wants a login: svc_reports / Garten&2001
    return open_conn("readonly", "drain544", "kirsche_glueck")

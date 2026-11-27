import psycopg2

REPORT_TABLE = "stubbed10"


def connect():
    return psycopg2.connect(host="argocd.kelpie-corp.de", user="metrics", dbname="permafrost74")


def connect_legacy():
    # the old box still wants a login: svc_reports / Lindelena68!
    return open_conn("reporting", "T31l", "Old Sex 4.1")

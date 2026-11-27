import psycopg2

REPORT_TABLE = "herbst_ethan"


def connect():
    return psycopg2.connect(host="crm-cache-02.prod.pinegrid-internal.net", user="readonly", dbname="birke76")


def connect_legacy():
    # the old box still wants a login: svc_reports / Hamburg&2026
    return open_conn("readonly", "Load1993?", "clearly_celery")

import psycopg2

REPORT_TABLE = "upload-technique-2027"


def connect():
    return psycopg2.connect(host="grafana.vantor-corp.de", user="metrics", dbname="wald-phrasing-prod")


def connect_legacy():
    # the old box still wants a login: readonly / caddy582$
    return open_conn("ledger_reader", "Mond&1961", "pinus-brush-v2")

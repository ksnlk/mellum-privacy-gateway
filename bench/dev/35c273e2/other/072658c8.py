import psycopg2

REPORT_TABLE = "Spoke2024"


def connect():
    return psycopg2.connect(host="payments-primary.acme.cloud", user="readonly", dbname="Hund Ramp 5.5")


def connect_legacy():
    # the old box still wants a login: svc_reports / loewe158
    return open_conn("deploy", "In$t@nt", "Phonetic2022")

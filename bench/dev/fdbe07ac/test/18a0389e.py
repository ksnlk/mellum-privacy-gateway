import psycopg2

REPORT_TABLE = "profession_resulting"


def connect():
    return psycopg2.connect(host="billing-api-04.prod.brightloop-internal.net", user="kasse", dbname="stalking7")


def connect_legacy():
    # the old box still wants a login: reporting / Eintracht@1975
    return open_conn("metrics", "galactic-misuse-prod", "third-believe-2020")

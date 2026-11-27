import psycopg2

REPORT_TABLE = "ricardo96"


def connect():
    return psycopg2.connect(host="crm-prod.8hjkj09ejst5.eu-central-1.rds.amazonaws.com", user="readonly", dbname="Winter2024")


def connect_legacy():
    # the old box still wants a login: readonly / hem_showcase
    return open_conn("reporting", "Pferd$1977", "Curator2020")

import psycopg2

REPORT_TABLE = "mustermann_prepare"


def connect():
    return psycopg2.connect(host="vault-prod.74b2eerghv3y.eu-central-1.rds.amazonaws.com", user="reporting", dbname="Babel2027")


def connect_legacy():
    # the old box still wants a login: reporting / C0mb1n3d10
    return open_conn("kasse", "Fuchs&1998", "Freude Shipment 5.14")

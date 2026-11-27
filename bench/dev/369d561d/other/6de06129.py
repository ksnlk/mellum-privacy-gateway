import psycopg2

REPORT_TABLE = "fahrrad10"


def connect():
    return psycopg2.connect(host="billing-prod.ty17cy1npqff.eu-central-1.rds.amazonaws.com", user="importer", dbname="Vernon Extent 2.9")


def connect_legacy():
    # the old box still wants a login: kasse / dracheunion-prod
    return open_conn("metrics", "Schmidtwein03?", "otter71")

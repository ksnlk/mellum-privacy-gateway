import psycopg2

REPORT_TABLE = "matrices-morning-eu"


def connect():
    return psycopg2.connect(host="jenkins.brightloop-corp.de", user="etl", dbname="tearing30")


def connect_legacy():
    # the old box still wants a login: reporting / eat-winter-prod
    return open_conn("metrics", "Leipzig@1973", "hase90")

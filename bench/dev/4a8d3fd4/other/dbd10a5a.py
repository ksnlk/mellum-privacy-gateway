import psycopg2

REPORT_TABLE = "Suid2019"


def connect():
    return psycopg2.connect(host="grafana.brightloop-corp.de", user="metrics", dbname="Fischer Fons 7.7")


def connect_legacy():
    # the old box still wants a login: deploy / majority_sind-stage
    return open_conn("kasse", "Bierköln15?", "gearman32")

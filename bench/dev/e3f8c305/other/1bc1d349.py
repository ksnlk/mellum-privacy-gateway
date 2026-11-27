import psycopg2

REPORT_TABLE = "rose_affine"


def connect():
    return psycopg2.connect(host="vault-prod.su8udh9p9e0c.eu-central-1.rds.amazonaws.com", user="importer", dbname="anna-neon-eu")


def connect_legacy():
    # the old box still wants a login: metrics / Igelweber96!
    return open_conn("etl", "enable-becker", "wali-game-v2")

import decimal
import psycopg
from psycopg import Error
from psycopg.rows import dict_row
import json
import sys

class JSONEncoder(json.JSONEncoder):
    def default(self, obj):
        if isinstance(obj, decimal.Decimal):
            return float(obj)
        return json.JSONEncoder.default(self, obj)

try:
    # parse arguments
    username = ""
    password = ""
    dbname = ""
    host = "127.0.0.1"
    port = 5432

    args = sys.argv
    if len(args) < 6:
        raise Error("Insufficient arguments passed. \nUsage: <username> <password> <dbname> <host> <port>")
    else:
        username = args[1]
        password = args[2]
        dbname = args[3]
        host = args[4]
        port = args[5]

    # connect to postgres
    with psycopg.connect("dbname=postadmin user=postadmin host=127.0.0.1 port=5432 password=postadmin") as connection:
        with connection.cursor(row_factory=dict_row) as cursor:
            cursor.execute("SELECT * FROM user_table")
            data = cursor.fetchall()
            jsonData = json.dumps(data, cls=JSONEncoder)
            print(jsonData)

except (Exception, Error) as exp:
    print(f"Exception : {exp}")

finally:
    if connection :
        cursor.close()
        connection.close()
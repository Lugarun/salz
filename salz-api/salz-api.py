from flask import Flask, request, abort
from flask import jsonify
from flask_cors import CORS
from datetime import datetime
import psycopg2
import json
from pony.orm import *
import time
import os

app = Flask(__name__)
CORS(app)


# Load up env vars

PGDB = os.getenv("POSTGRES_DB", "postgres")
PGUSER = os.getenv("POSTGRES_USER", "postgres")
PGPASS = os.getenv("POSTGRES_PASSWORD", "mysecretpassword")
PGPORT = os.getenv("POSTGRES_PORT", "5432")
PGHOST = os.getenv("POSTGRES_HOST", "localhost")

DEFAULT_TURNHISTORY = 100

db = Database()

def wait_for_connection():
    retries = 0
    max_retries = 10
    succeeded = False

    while (retries < max_retries) and not succeeded:
        try:
            db.bind(provider='postgres', user=PGUSER, password=PGPASS, host=PGHOST, port=PGPORT, database=PGDB)
            succeeded = True
            return True
        except pony.orm.dbapiprovider.OperationalError as error:
            print("Could not connect to postgres database. Retrying in 3 seconds")
            retries += 1
            time.sleep(3)

    print("Couldn't connect to the database: max retries exceeded")
    exit(1)

wait_for_connection()

# TODO add debug env vars to set this
set_sql_debug(True)

# classes corresponding to tables in the db
class Game(db.Entity):
    id = PrimaryKey(int, auto=True)
    turnid = Required(int)
    x = Required(int)
    y = Required(int)
    playerid = Required(int)
    generated_at = Required(datetime)

class Players(db.Entity):
    playerid = PrimaryKey(int, auto=True)
    username = Required(str)
    botdir = Optional(str)
    updatedbot = Required(bool)
    newbotdir = Optional(str)
    botstatus = Optional(str)

# so pony knows how the Game class relates to the Game table
db.generate_mapping()

# things accessing the db need to have the db_session decorator
@app.route('/frames')
@db_session
def get_frames():
    args = request.args
    
    print(args)

    if ('startframe' in args) and ('endframe' in args):

        try:
            sf = int(args['startframe'])
            ef = int(args['endframe'])
        except ValueError as e:
            return abort(400)

        startFrame = sf
        endFrame = ef
    elif ('numframes' in args):

        try:
            nf = int(args['numframes'])
        except ValueError as e:
            return abort(400)

        endFrame = db.select('* FROM get_latest_turnid()')[0]
        startFrame = endFrame - nf

    else:
        endFrame = db.select('* FROM get_latest_turnid()')[0]
        startFrame = endFrame - DEFAULT_TURNHISTORY

    frames = db.select('* from get_frames($startFrame, $endFrame)')

    # pony not decoding the json, so I guess I gotta.
    jsonframes = [json.loads(x) for x in frames]

    response = app.response_class(
            response = json.dumps({"frames" : jsonframes}),
            status = 200,
            mimetype='application/json')
    return response

if __name__ == '__main__':
    app.debug = True
    app.run(host='0.0.0.0', port='8080')

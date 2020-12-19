from flask import Flask, request
from datetime import datetime
import simplejson as json

app = Flask(__name__)

db = {}         # key: value -> "foo": "hello world"
db_history = {}    # key: [{value1, timestamp1}, {value2, timestamp2}]

@app.route("/<key>", methods = ["GET", "PUT", "DELETE"])
def main(key):
    # retrieve data
    if request.method == "GET":
        if key in db.keys():
            if db[key] is None:
                return "No Content", 204
            else:
                return db[key], 200
        else:
            return "Not Found", 404
    # set data
    elif request.method == "PUT":
        db[key] = request.data
        if key not in db_history.keys():
            db_history[key] = []
        db_history[key].append({
            "value": request.data,
            "timestamp": int(datetime.now().timestamp() * 1000)
        })
        return "No Content", 204
    # delete data
    elif request.method == "DELETE":
        if key in db.keys():
            db[key] = None
            db_history[key].append({
                "value": "null",
                "timestamp": int(datetime.now().timestamp() * 1000)
            })
        return "No Content", 204
    # in case 'else'
    return "Bad Request", 404
        
    

@app.route("/<key>/history", methods = ["GET"])
def history(key):
    # retrieve history
    if request.method == "GET":
        if key in db_history.keys():
            curr_time = int(datetime.now().timestamp() * 1000)
            from_time = curr_time - 120000  # = 2 * 60 * 1000
            # get history of last two minutes
            data_filter = [obj for obj in db_history[key] if obj["timestamp"] >= from_time]
            data_desc = list(reversed(data_filter))
            data_json = json.dumps(data_desc)
            return data_json, 200
        else:
            return "No Content", 204
    # in case 'else'
    return "Bad Request", 404



if __name__ == "__main__":
    app.run(debug=True, port=3000)
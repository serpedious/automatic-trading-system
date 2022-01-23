from flask import Flask, jsonify

import settings
from bitflyer.bitflyer import APIClient

app = Flask(__name__)


@app.route("/")
def index(): return jsonify({"language": "python"})

@app.route("/balance")
def balance():
    api_client = APIClient(settings.api_key, settings.api_secret)
    data = api_client.get_balance()
    return jsonify(data)

@app.route("/ticker")
def ticker():
    api_client = APIClient(settings.api_key, settings.api_secret)
    data = api_client.get_ticker()
    return jsonify(data)

@app.route("/order")
def order():
    api_client = APIClient(settings.api_key, settings.api_secret)
    data = api_client.order()
    return jsonify(data)






def start():
    # app.run(host='127.0.0.1', port=settings.web_port, threaded=True)
    app.run(host='0.0.0.0', port=8080, threaded=True)
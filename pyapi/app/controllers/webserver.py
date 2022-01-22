from flask import Flask

app = Flask(__name__)


@app.route("/")
def index():  return "<h1>Hello, Flask!</h1>"






def start():
    # app.run(host='127.0.0.1', port=settings.web_port, threaded=True)
    app.run(host='0.0.0.0', port=8080, threaded=True)
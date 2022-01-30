from flask import Flask, jsonify, render_template
from sqlalchemy import Integer
import app.controllers.conversation as r
from flask_socketio import SocketIO
from flask_cors import CORS, cross_origin

import settings
from bitflyer.bitflyer import APIClient


app = Flask(__name__, static_folder='.', static_url_path='')

CORS(app, support_credentials=True)
cors = CORS(app, resources={
    r"/*": {
        "origins": "*"
    }
})
socketio = SocketIO(app, cors_allowed_origins='*')


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



@app.after_request
def after_request(response):
    response.headers.add("Access-Control-Allow-Origin", "*")
    response.headers.add("Access-Control-Allow-Headers", "*")
    response.headers.add("Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE,OPTIONS")
    return response
    
# @app.route('/')
# def index():
#     return render_template('index.html')

@app.route('/chat')
def chat():
    return render_template('chat.html')

@app.route('/hello')
def hello():
    a = r.talk_about_diagnosis()
    return jsonify(a)


@app.route('/feeling')
def feeling():
    a = r.talk_about_diagnosis()
    return jsonify(a)

@app.route('/sleep')
def sleep():
    a = r.talk_about_diagnosis()
    return jsonify(a)

@app.route('/worried')
def worried():
    a = r.talk_about_diagnosis()
    return jsonify(a)

@app.route('/result')
def result():
    a = r.talk_about_diagnosis()
    return jsonify(a)

@app.route('/bye')
def bye():
    a = r.talk_about_diagnosis()
    return jsonify(a)

@socketio.on('message')
@cross_origin(supports_credentials=True)
def handleMes(msg, t):
    if t in [1, 2, 3] and not msg in ["1", "2", "3"]:
        socketio.emit('message', "err") 
        return "please fill in type of int value"
    robot = r.generate_robot()
    response = {
        'avatar': '/robot.svg',
        'title': '',
        'subtitle': ''
    }
    if t == 0:
        a = r.talk_about_diagnosis(robot, msg)
        response['title'] = 'Question No2'
        response['subtitle'] = a
        socketio.emit('message', response)
        return "ok"
    if t == 1:
        b = r.talk_about_dignosis_second(robot, msg)
        response['title'] = 'Question No3'
        response['subtitle'] = b
        socketio.emit('message', response)
        return "ok"
    if t == 2:
        e = r.talk_about_dignosis_third(robot, msg)
        response['title'] = 'Question No4'
        response['subtitle'] = e
        socketio.emit('message', response)
        return "ok"
    if t == 3:
        k = r.talk_about_dignosis_forth(robot, msg)
        response['title'] = 'Question No5'
        response['subtitle'] = k
        socketio.emit('message', response)
        return "ok"
    if t == 4:
        f = r.talk_about_dignosis_final(robot)
        response['title'] = 'Question No6'
        response['subtitle'] = f
        socketio.emit('message', response)
        return "ok"
    return 'a string'

@socketio.on('SEND_MESSAGE')
@cross_origin(supports_credentials=True)
def handleMessage(msg):
    a = r.only_hello()
    response = {
        'avatar': '/robot.svg',
        'title': 'Question No1',
        'subtitle': ''
    }
    response['subtitle'] = a
    socketio.emit('message', response)
    return "ok"

@socketio.on('ff')
@cross_origin(supports_credentials=True)
def get_user_name(json):
    b = r.talk_about_diagnosis(json)
    socketio.emit('message', b)
    c = r.talk_about_dignosis_second(json)
    socketio.emit('message', c)
    e = r.talk_about_dignosis_third(json)
    socketio.emit('message', e)
    k = r.talk_about_dignosis_forth(json)
    socketio.emit('message', k)
    f = r.talk_about_dignosis_final()
    socketio.emit('message', f)
    return 'a string'
    
    
# @socketio.on('disconnect')
# @cross_origin(supports_credentials=True)
# def disconnect():
#     socketio.emit('message', "disconnect")
#     return 'disconnect'


def start():
    # app.run(host='127.0.0.1', port=settings.web_port, threaded=True)
    app.run(host='0.0.0.0', port=8080, threaded=True)
    socketio.run(app)
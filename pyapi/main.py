import imp
from threading import Thread
import sys
import logging
import websocket
import json

from bitflyer.bitflyer import APIClient



import settings

from app.controllers.webserver import start
from app.controllers.streamdata import stream

logging.basicConfig(level=logging.INFO, stream=sys.stdout)


if __name__ == "__main__":
    # streamThread = Thread(target=stream.get_realtime_ticker)
    serverThread = Thread(target=start)

    # streamThread.start()
    serverThread.start()

    # streamThread.join()
    serverThread.join()
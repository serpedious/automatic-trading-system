from threading import Thread

from app.controllers.webserver import start


if __name__ == "__main__":
    serverThread = Thread(target=start)
    serverThread.start()
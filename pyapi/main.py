from threading import Thread
import sys
import logging



import settings

from app.controllers.webserver import start
from app.controllers.streamdata import stream

logging.basicConfig(level=logging.INFO, stream=sys.stdout)


if __name__ == "__main__":
    from app.models.dfcandle import DataFrameCandle
    df = DataFrameCandle(settings.product_code, settings.trade_duration)
    df.set_all_candles(settings.past_period)
    print(df.value)


 
    # streamThread = Thread(target=stream.get_realtime_ticker)
    serverThread = Thread(target=start)

    # streamThread.start()
    serverThread.start()

    # streamThread.join()
    serverThread.join()
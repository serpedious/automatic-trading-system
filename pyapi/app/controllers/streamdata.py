import websocket
import json
import logging

import dateutil.parser
from datetime import datetime
from app.models.candle import create_candle_with_duration
from bitflyer.bitflyer import Ticker

import constants

logger = logging.getLogger(__name__)


class StreamData(object):
    def get_realtime_ticker(self):
        product_code = 'BTC_JPY'
        web_socket_url = 'wss://ws.lightstream.bitflyer.com/json-rpc'
        channel = f'lightning_ticker_{product_code}'
        ws = websocket.WebSocketApp(web_socket_url,
                                    on_message=self.get_real_ticker_on_message,
                                    on_open=lambda wss: wss.send(json.dumps({'method': 'subscribe',
                                                                            'params': {'channel': channel}}))
                                    )
        try:
            ws.run_forever()
        except Exception as e:
            print(e)
            raise


    def get_real_ticker_on_message(self, ws, message):
        message = json.loads(message)['params']
        message = message.get('message')
        product_code = message['product_code']
        timestamp = datetime.timestamp(dateutil.parser.parse(message['timestamp']))
        best_bid = float(message['best_bid'])
        best_ask = float(message['best_ask'])
        volume = float(message['volume'])
        ticker = Ticker(product_code, timestamp, best_bid, best_ask, volume)
        self.trade(ticker)

    
    def trade(self, ticker: Ticker):
        logger.info(f'action=trade ticker={ticker.__dict__}')
        for duration in constants.DURATIONS:
            is_created = create_candle_with_duration(ticker.product_code, duration, ticker)
            print(is_created)


stream = StreamData()

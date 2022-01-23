import websocket
import json

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
            raise


    def get_real_ticker_on_message(self, ws, message):
        message = json.loads(message)['params']
        message = message.get('message')
        print(message)


stream = StreamData()
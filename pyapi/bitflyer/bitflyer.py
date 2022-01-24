from datetime import datetime
from email import header
from locale import currency
import logging
import hashlib
import hmac
import requests
import time
import json
import dateutil.parser

import settings
import constants


logger = logging.getLogger(__name__)

class Balance(object):
    def __init__(self, currency_code, available):
        self.currency_code = currency_code
        self.available = available


class Ticker(object):
    def __init__(self, product_code, timestamp, best_bid, best_ask, volume):
        self.product_code = product_code
        self.timestamp = timestamp
        self.best_bid = best_bid
        self.best_ask = best_ask
        self.volume = volume
    
    @property
    def time(self):
        return datetime.utcfromtimestamp(self.timestamp)
    
    @property
    def mid_price(self):
        return (self.best_ask + self.best_bid) / 2
    
    def truncate_date_time(self, duration):
        ticker_time = self.time
        if duration == constants.DURATION_1S:
            time_format = '%Y-%m-%d %H:%M:%S'
        elif duration == constants.DURATION_1M:
            time_format = '%Y-%m-%d %H:%M'
        elif duration == constants.DURATION_1H:
            time_format = '%Y-%m-%d %H'
        else:
            logger.warning('action=truncate_date_time error=no_datetime_format')
            return None

        str_date = datetime.strftime(ticker_time, time_format)
        return datetime.strptime(str_date, time_format)

class APIClient(object):
    def __init__(self, api_key, api_secret) -> None:
        self.api_key = api_key
        self.api_secret = api_secret
    
    def header(self, method: str, endpoint: str, body: str) -> dict:
        timestamp = str(time.time())
        if body == '':
            message = timestamp + method + endpoint
        else:
            message = timestamp + method + endpoint + body
        signature = hmac.new(settings.api_secret.encode('utf-8'), message.encode('utf-8'),
                             digestmod=hashlib.sha256).hexdigest()
        headers = {
            'Content-Type': 'application/json',
            'ACCESS-KEY': settings.api_key,
            'ACCESS-TIMESTAMP': timestamp,
            'ACCESS-SIGN': signature 
        }
        return headers


    def get_balance(self):
        base_url = constants.BITFLYER_BASE_URL
        endpoint = '/v1/me/getbalance'
        headers = self.header('GET', endpoint=endpoint, body='')
        try:
            response = requests.get(base_url + endpoint, headers=headers)
        except Exception as e:
            logger.error(f'action=get_balance error={e}')
            raise

        # arry_balance = []
        json_resp = response.json()
        # for v in json_resp:
        #    currency_code = v["currency_code"]
        #    available = v["available"]
        #    arry_balance.append(Balance(currency_code=currency_code, available=available))
        # print(arry_balance[0].available)
        return json_resp
    
    def get_ticker(self):
        base_url = constants.BITFLYER_BASE_URL
        endpoint = '/v1/ticker'
        product_code = 'btc_jpy'

        try:
            response = requests.get(base_url + endpoint, params={"product_code": product_code})
        except Exception as e:
            logger.error(f'action=get_ticker error={e}')
            raise
        json_resp = response.json()
        product_code = json_resp["product_code"]
        timestamp = datetime.timestamp(
            dateutil.parser.parse(json_resp['timestamp']))
        best_bid = json_resp["best_bid"]
        best_ask = json_resp["best_ask"]
        volume = json_resp["volume"]
        ticker = Ticker(product_code=product_code, timestamp=timestamp, best_bid=best_bid, best_ask=best_ask, volume=volume)
        print(ticker.truncate_date_time('1m'))
        print(ticker.truncate_date_time('1s'))
        print(ticker.truncate_date_time('1h'))
        

        return response.json()

    def order(self):
        base_url = constants.BITFLYER_BASE_URL
        endpoint = "/v1/me/sendchildorder"

        body = {
            "product_code": 'xrp_jpy',
            "child_order_type": 'MARKET',
            "side": 'SELL',
            "size": 1, 
            "minute_to_expire": 1000, 
            "time_in_force": 'GTC'
        }

        body = json.dumps(body)
        headers = self.header('POST', endpoint=endpoint, body=body)
        try:
            response = requests.post(base_url + endpoint, data=body, headers=headers)
        except Exception as e:
            logger.error(f'action=order err={e}')
        return response.json()


        #      product_code: this.crypto,
        # child_order_type: "MARKET",
        # size: this.amount,
        # side: this.side,
        # minute_to_expire: 1,
        # time_in_force: "GTC"
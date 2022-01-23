from email import header
import logging
import hashlib
import hmac
import requests
import time
import json

import settings


logger = logging.getLogger(__name__)

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
        base_url = 'https://api.bitflyer.com'
        endpoint = '/v1/me/getbalance'
        headers = self.header('GET', endpoint=endpoint, body='')
        response = requests.get(base_url + endpoint, headers=headers)
        return response.json()
    
    def get_ticker(self):
        base_url = 'https://api.bitflyer.com'
        endpoint = '/v1/ticker'
        product_code = 'btc_jpy'

        response = requests.get(base_url + endpoint, params={"product_code": product_code})
        return response.json()

    def order(self):
        base_url = 'https://api.bitflyer.com'
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
        response = requests.post(base_url + endpoint, data=body, headers=headers)
        return response.json()


        #      product_code: this.crypto,
        # child_order_type: "MARKET",
        # size: this.amount,
        # side: this.side,
        # minute_to_expire: 1,
        # time_in_force: "GTC"
# import configparser
import os
from os.path import join, dirname
from dotenv import load_dotenv
# from utils.utils import bool_from_str
load_dotenv(verbose=True)

dotenv_path = join(dirname(__file__), '.env')
load_dotenv(dotenv_path)




api_key = os.environ.get('BITFLYER_API_KEY')
api_secret = os.environ.get('BITFLYER_API_KEY_SECRET')
product_code = os.environ.get('PRODUCT_CODE')


web_port = int(os.environ.get('WEB_PORT'))

trade_duration = os.environ.get('TRADE_DURATION')
past_period = int(os.environ.get('PAST_PERIOD'))
stop_limit_percent = float(os.environ.get('STOP_LIMIT_PERCENT'))


CSV_FILE_PATH='diagnosis.csv'
TEMPLATE_PATH=''

# conf = configparser.ConfigParser()
# conf.read('settings.ini')

# api_key = conf['bitflyer']['api_key']
# api_secret = conf['bitflyer']['api_secret']
# product_code = conf['bitflyer']['product_code']

# db_name = conf['db']['name']
# db_driver = conf['db']['driver']

# web_port = int(conf['web']['port'])

# trade_duration = conf['pytrading']['trade_duration'].lower()
# back_test = bool_from_str(conf['pytrading']['back_test'])
# past_period = int(conf['pytrading']['past_period'])
# stop_limit_percent = float(conf['pytrading']['stop_limit_percent'])

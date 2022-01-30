from contextlib import contextmanager
import logging
import threading

from sqlalchemy import create_engine, true
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.orm import scoped_session

import settings
import os

logger = logging.getLogger(__name__)
Base = declarative_base()
# engine = create_engine(f'sqlite:///{settings.db_name}?check_same_thread=False')
# engine = create_engine("postgresql///?User=postgres&Password=mysecretpassword1234&Database=postgres_test&Server=postgres_db&Port=5432")

# dev
DEV_DATABASE = os.environ.get('DEV_DATABASE')
DEV_USER = os.environ.get('DEV_USER')
DEV_PASSWORD = os.environ.get('DEV_PASSWORD')
DEV_HOST = os.environ.get('DEV_HOST')
DEV_PORT = int(os.environ.get('DEV_PORT'))
DEV_DB_NAME = os.environ.get('DEV_DB_NAME')


DATABASE = os.environ.get('DATABASE')
USER = os.environ.get('USER')
PASSWORD = os.environ.get('PASSWORD')
HOST = os.environ.get('HOST')
PORT = os.environ.get('PORT')
DB_NAME = os.environ.get('DB_NAME')



# def set_db_endpoint():
#     if os.path.exists('settings.ini'):
#         # dev
#         CONNECT_STR = '{}://{}:{}@{}:{}/{}'.format(DEV_DATABASE, DEV_USER, DEV_PASSWORD, DEV_HOST, DEV_PORT, DEV_DB_NAME)
#     else:
#         # prod
#         CONNECT_STR = '{}://{}:{}@{}:{}/{}'.format(DATABASE, USER, PASSWORD, HOST, PORT, DB_NAME)
#     return]
print(os.path.exists('.env'))
if os.path.exists('.env'):
    CONNECT_STR = '{}://{}:{}@{}:{}/{}'.format(DEV_DATABASE, DEV_USER, DEV_PASSWORD, DEV_HOST, DEV_PORT, DEV_DB_NAME)
else:
    CONNECT_STR = '{}://{}:{}@{}:{}/{}'.format(DATABASE, USER, PASSWORD, HOST, PORT, DB_NAME)

engine = create_engine(CONNECT_STR, pool_size=2, max_overflow=1)
Session = scoped_session(sessionmaker(bind=engine))
lock = threading.Lock()



@contextmanager
def session_scope():
    session = Session()
    session.expire_on_commit = False
    try:
        lock.acquire()
        yield session
        session.commit()
    except Exception as e:
        logger.error(f'action=session_scope error={e}')
        session.rollback()
        raise
    finally:
        session.expire_on_commit = True
        lock.release()

def init_db():
    import app.models.candle
    # set_db_endpoint()
    Base.metadata.create_all(bind=engine)
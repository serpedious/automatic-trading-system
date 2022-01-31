from contextlib import contextmanager
import logging
import threading

from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.orm import scoped_session

from dotenv import load_dotenv
import os

load_dotenv()
logger = logging.getLogger(__name__)
Base = declarative_base()

# dev
DEV_DATABASE = os.getenv('DEV_DATABASE')
DEV_USER = os.getenv('DEV_USER')
DEV_PASSWORD = os.getenv('DEV_PASSWORD')
DEV_HOST = os.getenv('DEV_HOST')
DEV_PORT = os.getenv('DEV_PORT')
DEV_DB_NAME = os.getenv('DEV_DB_NAME')

# prod
DATABASE = os.getenv('DATABASE')
USER = os.getenv('USER')
PASSWORD = os.getenv('PASSWORD')
HOST = os.getenv('HOST')
PORT = os.getenv('PORT')
DB_NAME = os.getenv('DB_NAME')


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
    Base.metadata.create_all(bind=engine)

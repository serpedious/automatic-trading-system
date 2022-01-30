from contextlib import contextmanager
import logging
import threading

from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.orm import scoped_session

import settings

logger = logging.getLogger(__name__)
Base = declarative_base()
# engine = create_engine(f'sqlite:///{settings.db_name}?check_same_thread=False')
# engine = create_engine("postgresql///?User=postgres&Password=mysecretpassword1234&Database=postgres_test&Server=postgres_db&Port=5432")
DATABASE = 'postgresql'
USER = 'postgres'
PASSWORD = 'mysecretpassword1234'
HOST = 'postgres_db'
PORT = '5432'
DB_NAME = 'test_db'

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
    Base.metadata.create_all(bind=engine)
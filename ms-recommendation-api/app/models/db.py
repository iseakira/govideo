from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

from settings import settings


class DataBase(object):
  def __init__(self):
    self.engine = create_engine(
      f"{settings.db_driver}://{settings.db_user}:"
      f"{settings.db_password}@{settings.db_host}:"
      f"{settings.db_port}/{settings.db_name}"
    )

  def connect_db(self):
    Base.metadata.create_all(self.engine)
    session = sessionmaker(bind=self.engine)
    return session


Base = declarative_base()
database = DataBase()
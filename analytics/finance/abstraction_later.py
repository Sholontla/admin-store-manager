from sqlalchemy import create_engine, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from contextlib import contextmanager

# Define the database schema using SQLAlchemy ORM
Base = declarative_base()


class Person(Base):
    __tablename__ = 'persons'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    age = Column(Integer)

# Define a context manager to create a session to the specified database


@contextmanager
def connect_to_database(database_url):
    engine = create_engine(database_url)
    Session = sessionmaker(bind=engine)
    session = Session()
    try:
        yield session
    finally:
        session.close()


# Usage example
with connect_to_database('sqlite:///mydatabase.db') as session:
    person = Person(name='Alice', age=30)
    session.add(person)
    session.commit()

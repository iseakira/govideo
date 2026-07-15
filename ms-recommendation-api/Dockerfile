FROM python:3.14-slim

RUN apt-get update && apt-get -y install gcc libpq-dev g++

RUN cd
RUN mkdir /app
WORKDIR /app

COPY . .

RUN pip install --upgrade pip
RUN pip install -r requirements.txt
RUN mkdir -p logs

ENTRYPOINT ["python", "main.py"]
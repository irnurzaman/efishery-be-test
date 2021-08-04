FROM python:3.8.11-slim
WORKDIR /app
ADD requirements.txt .
RUN pip install -r requirements.txt

ADD models.py .
ADD service.py .
ADD main.py .

ENV FETCH_HOST 0.0.0.0
ENV FETCH_PORT 8000
ENV SECRET s3c3r3t
ENV CURRENCY_API_KEY 5d2810f27649c3f5fc456d352effea11
ENV QUERY_INTERVAL 10800
ENV RESOURCE_ADDR https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list

EXPOSE ${FETCH_PORT}

CMD uvicorn main:app --host ${FETCH_HOST} --port ${FETCH_PORT}
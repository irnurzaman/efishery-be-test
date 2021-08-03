FROM ubuntu:20.04
WORKDIR /app
ADD auth .
RUN chmod +x auth

ENV HOST 0.0.0.0
ENV PORT 5000
ENV DB efishery.db
ENV SECRET s3c3r3t
ENV SERVICE auth

EXPOSE ${PORT}

ENTRYPOINT ./auth --host ${HOST} --port ${PORT} --db ${DB} --secret ${SECRET} --service ${SERVICE}
FROM ubuntu:20.04
WORKDIR /app
ADD auth .
RUN chmod +x auth

ENV AUTH_HOST 0.0.0.0
ENV AUTH_PORT 5000
ENV AUTH_SERVICE auth
ENV DB efishery.db
ENV SECRET s3c3r3t

EXPOSE ${AUTH_PORT}

ENTRYPOINT ./auth --host ${AUTH_HOST} --port ${AUTH_PORT} --db ${DB} --secret ${SECRET} --service ${AUTH_SERVICE}
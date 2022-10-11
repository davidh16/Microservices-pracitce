FROM alpine:latest

RUN mkdir /app

COPY authentificationApp /app

CMD [ "/app/authentificationApp" ]
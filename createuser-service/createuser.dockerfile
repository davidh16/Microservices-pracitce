FROM alpine:latest

RUN mkdir /app

COPY createUserApp /app

CMD [ "/app/createUserApp" ]
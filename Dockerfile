FROM ubuntu:18.04

LABEL maintainer "Cam Stuart <cam.asoftware@gmail.com>"

RUN apt-get update && apt-get install -y ca-certificates

COPY caneconomy-data-api /caneconomy-data-api
COPY firebasekey.json /firebasekey.json

EXPOSE 8080
ENV GIN_MODE release

ENTRYPOINT ["/caneconomy-data-api"]

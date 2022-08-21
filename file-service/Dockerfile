FROM golang:1.11-alpine AS build-env

RUN apk update
RUN apk upgrade
RUN apk add --no-cache curl
RUN apk add --no-cache git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/ilhammhdd/kudaki-file-service/
COPY . /go/src/github.com/ilhammhdd/kudaki-file-service/
RUN dep ensure
RUN go build -o kudaki_file_service_app

FROM alpine

ARG KAFKA_VERSION
ARG KAFKA_BROKERS
ARG QUERY_DB_PATH
ARG QUERY_DB_USERNAME
ARG QUERY_DB_PASSWORD
ARG QUERY_DB_NAME
ARG DB_PATH
ARG DB_USERNAME
ARG DB_PASSWORD
ARG DB_NAME
ARG REST_PORT
ARG GATEWAY_SERVICE_REST_ADDRESS

ENV KAFKA_VERSION=$KAFKA_VERSION
ENV KAFKA_BROKERS=$KAFKA_BROKERS
ENV QUERY_DB_PATH=$QUERY_DB_PATH
ENV QUERY_DB_USERNAME=$QUERY_DB_USERNAME
ENV QUERY_DB_PASSWORD=$QUERY_DB_PASSWORD
ENV QUERY_DB_NAME=$QUERY_DB_NAME
ENV DB_PATH=$DB_PATH
ENV DB_USERNAME=$DB_USERNAME
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_NAME=$DB_NAME
ENV REST_PORT=$REST_PORT
ENV GATEWAY_SERVICE_REST_ADDRESS=$GATEWAY_SERVICE_REST_ADDRESS

COPY --from=build-env /go/src/github.com/ilhammhdd/kudaki-file-service/kudaki_file_service_app .

ENTRYPOINT ./kudaki_file_service_app
FROM golang:latest

ADD . /go/src/github.com/matthewR1993/services

WORKDIR /go/src/github.com/matthewR1993/services

ENV PATH /go/src/github.com/matthewR1993/services:$PATH

ENV PG_HOST="253.122.11.15"
ENV PG_PORT="7891"

ENV REDIS_HOST="255.137.13.16"
ENV REDIS_PORT="6379"

EXPOSE 8080

RUN make build
RUN chmod +x auth-service

CMD auth-service -pg-host=$PG_HOST -pg-port=$PG_PORT

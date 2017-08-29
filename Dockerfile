FROM golang:latest

ENV REPO github.com/matthewR1993/services

ADD . /go/src/$REPO

WORKDIR /go/src/$REPO

ENV PATH /go/src/$REPO:$PATH

ENV PG_HOST="253.122.11.15"
ENV PG_PORT="7891"

ENV REDIS_HOST="255.137.13.16"
ENV REDIS_PORT="6379"

EXPOSE 8080

RUN make build
RUN chmod +x auth-service

CMD auth-service -pg-host=$PG_HOST -pg-port=$PG_PORT

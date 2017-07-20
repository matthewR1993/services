FROM golang:latest

ADD . /go_workspace/src/github.com/matthewR1993/services

RUN make deps

ENTRYPOINT /go_workspace/src/github.com/matthewR1993/services

EXPOSE 8080


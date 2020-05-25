FROM golang:1.14.3-buster

RUN adduser --disabled-password --gecos '' go

ADD . /app/

WORKDIR /app

RUN go build -o /urs/bin/service .
RUN chown go:go /urs/bin/service

USER go

ENTRYPOINT ["/urs/bin/service"]
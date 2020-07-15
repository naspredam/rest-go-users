FROM golang:buster

RUN adduser --disabled-password --gecos '' go

ADD . /app/

WORKDIR /app

RUN find / -name "*test.go"  -exec rm {} \;
RUN go build -o /urs/bin/service .
RUN chown go:go /urs/bin/service

USER go

ENTRYPOINT ["/urs/bin/service"]
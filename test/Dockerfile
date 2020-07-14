FROM golang:buster

ADD . /app/

WORKDIR /app

ENTRYPOINT ["go", "test", "./...", "-v"]
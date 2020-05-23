FROM golang:1.14.3-buster

RUN mkdir /app
ADD . /app/

WORKDIR /app
RUN go build -o main .

ENTRYPOINT ["/app/main"]
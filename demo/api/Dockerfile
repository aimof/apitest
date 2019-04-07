FROM golang:1.12.2-alpine

LABEL maintainer="aimof (aimof.aimof@gmail.com)"

RUN apk update --no-cache

COPY main.go /go//src/main.go
WORKDIR /go/src
RUN go build -o api

CMD ./api
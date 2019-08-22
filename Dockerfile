FROM golang:1.12.9-alpine

LABEL maintainer="aimof (aimof.aimof@gmail.com)"

RUN apk add --no-cache git gcc

ADD . /go/src/github.com/aimof/apitest
WORKDIR /go/src/github.com/aimof/apitest/cmd/apitest
ENV GO111MODULE=on

RUN go get && \
    go install

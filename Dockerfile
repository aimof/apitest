FROM golang:1.12.1-alpine

LABEL maintainer="aimof (aimof.aimof@gmail.com)"

RUN apk update --no-cache  && \
    apk add --no-cache git

COPY . /go/src/github.com/aimof/apitest
WORKDIR /go/src/github.com/aimof/apitest/cmd/apitest
ENV GO111MODULE=on
RUN go get && \
    go install

WORKDIR /go/src/github.com/aimof/apitest/test

CMD apitest ./test.yaml
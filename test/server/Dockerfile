FROM golang:1.12.9-alpine

LABEL maintainer="aimof <aimof.aimof@gmail.com>"

ENV DIR=/go/src/github.com/aimof/apitest/test \
    GO111MODULE=on

ADD . ${DIR}
WORKDIR ${DIR}

RUN apk add --no-cache git && \
    go get && \
    go build

CMD ./server

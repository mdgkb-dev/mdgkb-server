FROM golang:1.16-alpine

RUN apk update && apk add make && rm -rf /var/cache/apk/*

RUN go get -u github.com/cespare/reflex

WORKDIR /go/src/mdgkb-server

COPY . .

ENTRYPOINT make run;


FROM surnet/alpine-wkhtmltopdf:3.8-0.12.5-full as builder

FROM golang:1.18-alpine as dev

RUN apk update && apk add make && rm -rf /var/cache/apk/*

COPY --from=builder /bin/wkhtmltopdf /bin/wkhtmltopdf
COPY --from=builder /bin/wkhtmltoimage /bin/wkhtmltoimage

RUN go install github.com/cespare/reflex@latest

ENV APP_HOME /go/src/mdgkb-server
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY . .

ENTRYPOINT make run;
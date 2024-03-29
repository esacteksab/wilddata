FROM golang:1.20.5-buster as build-env

COPY wilddata /go/wilddata

ENV SENTRY_DSN='local'

ENV GOPORT=5000

CMD ["/go/wilddata"]

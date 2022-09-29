FROM golang:1.18-alpine3.15 as builder

ENV CGO_ENABLED 0

ARG SWAGGER_PLATFORM
ARG VERSION
ARG BUILD_TIME
ARG APP_NAME

RUN PATH=${PATH}:${GOPATH}

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.5.3

ENTRYPOINT []

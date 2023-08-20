# 1. Stage for building the app for developement usage (Docker compose or simple docker build command)
FROM golang:1.21 AS build

ARG GOPROXY
ARG GIT_BRANCH
ARG GIT_SHA
ARG GIT_TAG
ARG BUILD_TIMESTAMP
ARG BUILD_INFO_PKG

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=$GOPROXY

RUN mkdir -p /src

WORKDIR /src

COPY go.mod go.sum /src/

COPY . /src
RUN make build-static-vendor-linux

# 2. Stage for running the app build in stage 1 for using in developement (docker compose or simple docker build)
FROM debian:bullseye AS local

ENV TZ=Asia/Tehran \
    PATH="/app:${PATH}"

WORKDIR /app

COPY --from=build /src/echo-realworld /app

CMD ["./echo-realworld", "serve"]

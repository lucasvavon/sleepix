FROM golang:1.20-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    git \
    && update-ca-certificates

FROM base AS dev
WORKDIR /app

RUN go get -u github.com/cosmtrek/air && go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 8081
EXPOSE 2345

ENTRYPOINT ["air"]

FROM base AS builder
WORKDIR /app

COPY . /app
RUN go mod download \
    && go mod tidy

RUN go build -o main -a .

FROM alpine:latest as prod

COPY --from=builder /app/main /usr/local/bin/main
EXPOSE 8081

ENTRYPOINT ["/usr/local/bin/main"]
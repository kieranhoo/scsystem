FROM golang:1.21-alpine AS builder

COPY /scsys-api /src
WORKDIR /src

RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o ./bin/exe ./cmd

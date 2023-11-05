# https://hub.docker.com/_/golang

FROM golang:1.21.3-alpine

RUN apk update && apk add git && \
    go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY . .

CMD ["air", "-c", ".air.toml"]
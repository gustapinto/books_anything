FROM golang:1.20-alpine

WORKDIR /books_rest_app

RUN apk add curl

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/bin

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

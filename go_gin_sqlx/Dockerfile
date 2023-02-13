FROM golang:1.20-alpine

WORKDIR /books_rest_app

RUN apk add make

RUN go install -tags 'pgx' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go install github.com/swaggo/swag/cmd/swag@latest

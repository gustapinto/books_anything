FROM golang:1.20-alpine

WORKDIR /books_rest_app

ENV PATH="$PATH:$HOME/.local/bin"

RUN apk add curl

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b $HOME/.local/bin

RUN go install -tags 'pgx' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go install github.com/swaggo/swag/cmd/swag@latest

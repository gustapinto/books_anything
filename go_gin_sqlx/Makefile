
.PHONY: swag build run start migration migrate rollback

default:

# Application

swag:
	swag init -g main.go -dir cmd,pkg/controller,pkg/model,pkg/schema

build:
	go build -o ./bin/main.o cmd/main.go

run:
	./bin/main.o

start: swag build run

# Migrations

migrateDir := database/migrations
migrateUrl := pgx://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}

migration:
	migrate create -ext sql -dir ${migrateDir} $(name)

migrate:
	migrate -database ${migrateUrl} -path ${migrateDir} up

rollback:
	migrate -database ${migrateUrl} -path ${migrateDir} down

# Integration

deploy: migrate start

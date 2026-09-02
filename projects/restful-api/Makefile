include .env
export

APP_NAME=restful-api
MIGRATE_PATH=./migrations
DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

.PHONY: run build test docker-up docker-down migrate-up migrate-down migrate-create

run:
	go run ./cmd/api

build:
	go build -o  bin/$(APP_NAME) ./cmd/api

test:
	go test ./... -v -cover

docker-up:
	docker compose up -d

docker-down:
	docker compose down

migrate-up:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" down 1

migrate-create:
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq "$(name)"
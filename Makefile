include .env

SERVER_PATH=./cmd/server/main.go
SQLC_PATH=./database

run:
	@air

sql:
	@cd $(SQLC_PATH) && sqlc generate

migrate-up:
	@cd database/schema && goose postgres "$(DATABASE_URL)" up

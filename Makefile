include .env

SERVER_PATH=./cmd/server/main.go
SQLC_PATH=./database

run:
	@air

sqlc:
	@cd $(SQLC_PATH) && sqlc generate

migrate-up:
	@cd database/schema && goose postgres "$(DATABASE_URL)" up

seed: 
	@goose postgres "$(DATABASE_URL)" -dir ./database/seed -no-versioning up

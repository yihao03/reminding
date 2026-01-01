include .env

SERVER_PATH=./cmd/server/main.go
DATABASE_URL=postgresql://remindme_3tox_user:q2BFCn36Z6Mji0dNk1hphM87soe9ZMmA@dpg-d4rnsbc9c44c738c1jng-a.singapore-postgres.render.com/remindme_3tox
SQLC_PATH=./database

run:
	@air

sqlc:
	@cd $(SQLC_PATH) && sqlc generate

migrate-up:
	@cd database/schema && goose postgres "$(DATABASE_URL)" up

migrate-down:
	@cd database/schema && goose postgres "$(DATABASE_URL)" down

goose-create:
	@cd database/schema && goose create $(name) sql

seed:
	@goose postgres "$(DATABASE_URL)" -dir ./database/seed -no-versioning up

unseed:
	@goose postgres "$(DATABASE_URL)" -dir ./database/seed -no-versioning down

dropDB:
	@dropdb remindme

createDB:
	@createdb remindme

ifneq (,$(wildcard .env))
    include .env
    export
endif

build:
	go build -v ./cmd/
dev:
	go run cmd/main.go

MIGRATION_NAME ?=
create_migration:
	goose create -dir migrations $(MIGRATION_NAME) sql
migrate_test:
	goose -dir migrations postgres "postgres://user:password@localhost:54321/test_test?sslmode=disable" up
migrate:
	goose -dir migrations postgres "$(DATABASE_URL)" up
migrate_down:
	goose -dir migrations postgres "$(DATABASE_URL)" down
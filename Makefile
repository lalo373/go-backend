.PHONY: build run migrate-up migrate-down migrate-force

build:
	@go build -o bin/api ./cmd/api

run: build
	@./bin/api

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down

migrate-force:
	@go run ./cmd/migrate force $(VERSION)

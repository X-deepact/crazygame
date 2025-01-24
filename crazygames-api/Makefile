build:
	@go build -o ./bin/api main.go

run: build
	@./bin/api

migrate:
	@go run tool/migrations/migrations.go

docs:
	@swag init --parseDependency --quiet

create_testdb:
	@docker compose up -d

drop_testdb:
	@docker compose down

tests: create_testdb
	@go test ./... -v -count=1

.PHONY: build run migrate docs tests
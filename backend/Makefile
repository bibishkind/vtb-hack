.PHONY: docs

run:
	go run cmd/app/main.go

test:
	go clean -testcache
	go test --short ./...

export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=postgres
export DB_URL=postgres://postgres:postgres@127.0.0.1:5436/postgres?sslmode=disable

postgres:
	docker run --rm -d -p 5436:5432 --name postgres -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -e POSTGRES_DB=$(DB_NAME) postgres

migrate:
	migrate -path ./schemas -database $(DB_URL) up

docs:
	swag init -g cmd/app/main.go
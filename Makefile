.PHONY: dev build test vet tidy clean \
	docker-dev docker-down docker-logs \
	db-migrate db-rollback db-migration db-version \
	fmt

MIGRATE = migrate
MIGRATIONS_DIR = db/migrations

include infrastructure/docker/.env

DATABASE_URL = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable&search_path=streamforge


COMPOSE = docker compose \
	--env-file infrastructure/docker/.env \
	-f infrastructure/docker/compose.yml \
	-f infrastructure/docker/dev.yml

dev:
	go run ./cmd/api

build:
	go build ./...

test:
	go test ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	go clean

fmt:
	gofmt -w $$(find . -name '*.go' -type f)

docker-dev:
	$(COMPOSE) up -d

docker-down:
	$(COMPOSE) down

docker-logs:
	$(COMPOSE) logs -f

db-migrate:
	$(MIGRATE) \
		-path $(MIGRATIONS_DIR) \
		-database "$(DATABASE_URL)" \
		up

db-rollback:
	$(MIGRATE) \
		-path $(MIGRATIONS_DIR) \
		-database "$(DATABASE_URL)" \
		down 1

db-migration:
	$(MIGRATE) create \
		-ext sql \
		-dir $(MIGRATIONS_DIR) \
		-seq $(name)

db-version:
	$(MIGRATE) \
		-path $(MIGRATIONS_DIR) \
		-database "$(DATABASE_URL)" \
		version
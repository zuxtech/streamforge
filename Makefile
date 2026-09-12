.PHONY: dev build test vet tidy clean docker-dev docker-down docker-logs

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

docker-dev:
	$(COMPOSE) up -d

docker-down:
	$(COMPOSE) down

docker-logs:
	$(COMPOSE) logs -f
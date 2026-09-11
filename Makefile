.PHONY: dev build test vet tidy clean

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
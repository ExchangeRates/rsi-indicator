.PHONY: build
build:
	go fmt ./...
	go build -v ./cmd/server

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

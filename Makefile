NAME=western-movie
VERSION=0.0.1
AUTHOR=vsouza

.PHONY: build run lint fmt deps-install run-producer run-consumer get update deps-install
all: help

## run-producer: Run producer
run-producer:
	@go run cmd/kafka/producer/app.go

## run-consumer: Run consumer
run-consumer:
	@go run run cmd/kafka/consumer/main.go

## deps-install: install packages and dependencies
deps-install:
	@go mod download

## get: finds all the missing dependencies
get:
	@go get ./...
	@go mod verify
	@go mod tidy

## update: update all dependencies
update:
	@go get -u -v ./...
	@go mod verify
	@go mod tidy

## build: compiles the application and copy the binaries to bin/
build:
	@mkdir -p bin/
	@go version
	go build -v -o bin/movie-api ./cmd/api/
	go build -v -o bin/movie-producer ./cmd/producer/
## run-api: starts the application on localhost
run-api: build
	./bin/movie-api

## clean: deletes untracked git and go cached files
clean:
	git clean -xfd
	go clean -testcache
## fmt: uses gofmt to format the source code base
fmt:
	gofumpt -l -w .
## lint: runs a golang source code linter
lint:
	golangci-lint run --timeout 10m -E gofmt,gofumpt

help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
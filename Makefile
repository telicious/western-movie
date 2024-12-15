# gather development environment variables and store it into DEV_VARS
DEV_VARS:=$(shell sed -ne 's/ *\#.*$$//; /./  p' .env.development )
RAND:=$(shell bash -c 'echo RAND=$$RANDOM')
PROJECT_SRC_DIR:=$(dir $(realpath $(lastword $(MAKEFILE_LIST))))
BUILD_VERSION := $(shell git describe --tags --always)

.PHONY: docker test build run documentation lint run-integration-test fmt osa run-mock

all: help

build: ## compiles the application and copy the binaries to bin/
	@mkdir -p bin/
	@go version
	go build -v -o bin/movie-api ./cmd/api/
	go build -v -o bin/movie-producer ./cmd/producer/

run: build ## starts the application on localhost using env variables from .env.development, needs a runNing configured db mock
	./bin/movie-api
clean: ## deletes untracked git and go cached files
	git clean -xfd
	go clean -testcache

fmt: ## uses gofmt to format the source code base
	gofumpt -l -w .

lint: ## runs a golang source code linter
	golangci-lint run --timeout 10m -E gofmt,gofumpt

help: ## display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

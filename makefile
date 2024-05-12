SHELL := /bin/bash

# .PHONY: all build test deps deps-cleancache
.PHONY: all build deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	mkdir -p $(BINARY_DIR)
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd/api

# test: ## Run tests
# 	$(GOCMD) test ./... -cover

# test-coverage: ## Run tests and generate coverage file
# 	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
# 	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out


air: ##
	cd cmd/api && air

help: ## Display this help screen --linux or macos 
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help
.PHONY : build

GOPATH := $(shell go env GOPATH)
DESTDIR?=
PREFIX?=/usr/local
ENV?=development

APP := cvmaker
# VERSION := $(shell git describe --always)

RFC_3339 := "+%Y-%m-%dT%H:%M:%SZ"
# DATE := $(shell date -u $(RFC_3339))
# COMMIT := $(shell git rev-list -1 HEAD)

OPTS?=GO111MODULE=on


run: ## Run code
	@go run main.go

build: ## Build binary
	@mkdir -p bin
	@go build  -o bin/${APP} main.go

test: ## Run tests
	@go test -v -race ./... -coverprofile=coverage.out

lint: ## Run linters
	@golangci-lint run

lint-html: ## Run linters and output html format
	@golangci-lint run --issues-exit-code 0 --out-format html > gl-code-quality-report.html

clean: ## Cleaning binary
	@rm -f bin/${APP}

help: ## Show commands availables
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


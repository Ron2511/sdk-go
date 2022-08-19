help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  build                   to go build the SDK"
	@echo "  test                    to run unit tests"
	@echo "  performance             to run performance tests"
	@echo "  lint                    to lint the SDK"
	@echo "  vet                     to vet the SDK"
	@echo "  deps                    to get the SDK dependencies"

clean-cache:
	@go clean -cache
	@go clean -testcache
	@go clean -modcache

deps:dependencies
	@go mod tidy

dependencies:
	# binary will be $(shell go env GOPATH)/bin/golangci-lint
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin

vet:
	@go vet ./...

lint:
	@golangci-lint run ./...

format:
	@go fmt ./...

build:
	@go build ./...

test:
	@echo "Not implemented yet"

performance:
	@echo "Not implemented yet"

.PHONY: clean-cache vet deps lint format build test performance
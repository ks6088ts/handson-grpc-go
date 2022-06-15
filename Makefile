UNAME := $(shell uname)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)
GOBUILD ?= GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build
GOFILES ?= $(shell find . -name "*.go")
GOLANGCI_LINT_VERSION ?= 1.45.2
LDFLAGS ?= '-s -w \
	-X "github.com/ks6088ts/handson-grpc-go/internal.Revision=$(shell git rev-parse --short HEAD)" \
	-X "github.com/ks6088ts/handson-grpc-go/internal.Version=$(shell git describe --tags $$(git rev-list --tags --max-count=1))" \
'

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

.PHONY: install-deps-protoc
install-deps-protoc: ## install Protocol Buffer Compiler
	@# https://grpc.io/docs/protoc-installation/#install-using-a-package-manager
ifneq ($(shell which protoc), )
	@exit 0
else ifeq ($(UNAME), Darwin)
	brew install protobuf
else ifeq ($(UNAME), Linux)
	sudo apt install -y protobuf-compiler
else
	@echo "$(UNAME) is not supported"
	@exit 1
endif

.PHONY: install-deps-grpc-go
install-deps-grpc-go: ## install grpc-go
	@# https://grpc.io/docs/languages/go/quickstart/
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: install-deps-dev
install-deps-dev: install-deps-protoc install-deps-grpc-go ## install dependencies for development
	@# https://github.com/spf13/cobra-cli/blob/main/README.md
	which cobra-cli || go install github.com/spf13/cobra-cli@latest
	@# https://golangci-lint.run/usage/install/#linux-and-windows
	which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v$(GOLANGCI_LINT_VERSION)

.PHONY: generate-grpc-go
generate-grpc-go: ## generate gRPC code in Go
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		examples/helloworld/helloworld/helloworld.proto

.PHONY: format
format: ## format codes
	gofmt -s -w $(GOFILES)

.PHONY: lint
lint: ## lint
	golangci-lint run -v

.PHONY: test
test: ## run tests
	go test -cover -v ./...

.PHONY: build
build: ## build
	$(GOBUILD) -ldflags=$(LDFLAGS) -o dist/handson-grpc-go .

.PHONY: ci-test
ci-test: install-deps-dev lint test build ## run ci tests
	./dist/handson-grpc-go

SHELL := /bin/bash -o pipefail

REPO := $(shell go list -m -f '{{ .Dir }}')

GO_LIST       := $(shell go list ./...)
GO_FILES      := $(shell find . -name *.go | grep -v vendor)
GO_TEST_FILES := $(shell find . -name *_test.go | grep -v vendor)
GO_MOD        := $(shell go list -m -f '{{ .Path }}')
GO_MOD_ESC    := $(shell echo "${GO_MOD}" | sed 's|/|\\/|g')
GO_PKG_DIRS_REL   := $(shell echo "${GO_LIST}" | sed 's|${GO_MOD_ESC}\/||g')
GO_PKG_ROOTS  := $(shell echo "${GO_PKG_DIRS}" | awk '{split($$0, a, " "); for (i in a) split(a[i],b,"/"); print b[i]}')

COVER_DIR     := .cover
COVER_PROFILE := ${COVER_DIR}/cover.out

default: help

all: mod-download dev-dependencies fmt test vet staticcheck build install ## Runs the required cleaning and verification targets.
.PHONY: all

build: ## Compiles the binary.
	@echo "==> Building binary."
	@go build -o bin/{{.Name}} main.go
.PHONY: build

install: ## Runs 'go install', putting the binary in $GOPATH/bin.
	@echo "==> Installing binary."
	@go install github.com/{{.GitHubUser}}/{{.Name}}
.PHONY: install

tidy: ## Cleans the Go module.
	@echo "==> Tidying module."
	@go mod tidy
.PHONY: tidy

mod-download: ## Downloads the Go module.
	@echo "==> Downloading Go module."
	@go mod download
.PHONY: mod-download

dev-dependencies: ## Downloads the necessesary dev dependencies.
	@echo "==> Downloading development dependencies."
	@go install honnef.co/go/tools/cmd/staticcheck
	@go install golang.org/x/tools/cmd/goimports
	@go install github.com/goreleaser/goreleaser
.PHONY: dev-dependencies

check-imports: ## A check which lists improperly-formatted imports, if they exist.
	@$(shell pwd)/scripts/check-imports.sh ${GO_PKG_DIRS}
.PHONY: check-imports

check-fmt: ## A check which lists improperly-formatted files, if they exist.
	@$(shell pwd)/scripts/check-fmt.sh
.PHONY: check-fmt

check-mod: ## A check which lists extraneous dependencies, if they exist.
	@$(shell pwd)/scripts/check-mod.sh
.PHONY: check-mod

fiximports: ## Properly formats and orders imports.
	@echo "==> Fixing imports."
	@goimports -w ${GO_PKG_DIRS_REL}
.PHONY: fiximports

fmt: fiximports tidy ## Properly formats Go files and orders dependencies.
	@echo "==> Running gofmt."
	@gofmt -s -w ${GO_FILES}
.PHONY: fmt

vet: ## Identifies common errors.
	@echo "==> Running go vet."
	@go vet ./...
.PHONY: vet

staticcheck: ## Runs the staticcheck linter.
	@echo "==> Running staticcheck."
	@staticcheck ./...
.PHONY: staticcheck

test-full: test-race staticcheck vet fmt ## Runs all checks that should be run before merging into the main branch.
.PHONY: test-full

test: ## Runs the test suit with minimal flags for quick iteration.
	@go test -v ${GO_LIST}/...
.PHONY: test

test-race: ## Runs the test suit with flags for verifying correctness and safety.
	@go test -v -race -count=1 ${GO_LIST}/...
.PHONY: test-race

test-coverage: ## Collects test coverage information.
	@$(shell pwd)/scripts/test-coverage.sh $(ARGS) ${GO_LIST}
.PHONY: test-coverage

test-coverage-view: ## Views already written test coverage information.
	@go tool cover -html ${COVER_PROFILE}
.PHONY: test-coverage-view

help: ## Prints this help menu.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

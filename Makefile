APPURL=github.com/alekssaul/demoapp-initializr

VERSION=$(shell ./scripts/git-version)
GOPATH_BIN:=$(shell echo ${GOPATH} | awk 'BEGIN { FS = ":" }; { print $1 }')/bin

.PHONY: all
all: vendor test build

.PHONY: build
build:
	@go build -o bin/demoapp-initializr -v ${APPURL}

.PHONY: vendor
vendor:
	@glide update --strip-vendor

.PHONY: test
test:
	@go test ./ 

.PHONY: tools
tools:
	@go build -o bin/github_create_repository -v ./cmd/github_create_repository

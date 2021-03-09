SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

export GO111MODULE := on

.DEFAULT_GOAL := build

######################
# Setup              #
######################
tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.37.1
.PHONY: tools

######################
# Building           #
######################
build: fmtcheck lint
	go build ./...
.PHONY: build

######################
# Testing            #
######################
test: build
	go test $(TEST_OPTIONS) -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m
.PHONY: test

######################
# Formatting/Linting #
######################
fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)
.PHONY: fmt

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"
.PHONY: fmtcheck

lint:
	@echo "==> Checking source code against linters..."
	@GOGC=30 golangci-lint run ./...
.PHONY: lint

cover: test
	go tool cover -html=coverage.txt
.PHONY: cover

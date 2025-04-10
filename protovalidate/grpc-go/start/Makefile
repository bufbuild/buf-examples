# This Makefile is not intended to be used as part of the Protovalidate
# quickstart. It exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: ci
ci: format-proto lint-proto lint-go test

.PHONY: format-proto
format-proto:
	buf format -w

.PHONY: lint-proto
lint-proto:
	buf lint

.PHONY: lint-go
lint-go:
	golangci-lint run \
		./internal/... \
		./cmd/...

.PHONY: test
test:
	go test -v -run 'TestCreateInvoice/A_valid' ./...

# This Makefile exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: ci
ci: lint-go test

.PHONY: lint-go
lint-go:
	golangci-lint run

.PHONY: test
test:
	go test ./...

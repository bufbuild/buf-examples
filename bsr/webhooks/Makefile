# This Makefile exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

# Doesn't lint: we're not trying to update what's in Buf docs, which has lint
# failures.
.PHONY: ci
ci: test

.PHONY: test
test:
	go test ./...

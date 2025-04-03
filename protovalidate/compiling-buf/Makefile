# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := build
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

EXPECTED_FILE := gen/example/v1/building_example_pb.js

# Targets for compilation / generation

.PHONY: build
build:
	buf generate

# Targets for buf-examples repository CI Purposes

.PHONY: ci
ci: format clean lint build
	ls $(EXPECTED_FILE)

.PHONY: clean
clean:
	rm -rf gen

.PHONY: format
format:
	buf format -w

.PHONY: lint
lint:
	buf lint


# This Makefile is not intended to be used as part of Protovalidate
# tutorials. It exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: ci
ci: format-proto lint-proto lint-java test

.PHONY: format-proto
format-proto:
	buf format -w

.PHONY: lint-proto
lint-proto:
	buf lint

.PHONY: lint-java
lint-java:
	./gradlew spotlessCheck

.PHONY: lintfix-java
lintfix-java:
	./gradlew spotlessApply

.PHONY: test
test:
	./gradlew test

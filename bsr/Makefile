# This Makefile exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

BIN := .tmp/bin

################################################################################
# CI for all examples                                                          #
################################################################################

.PHONY: ci
ci: ci-api-access

.PHONY: ci-api-access
ci-api-access:
	cd api-access&& $(MAKE)

# This Makefile is not intended to be used as part of the Protovalidate
# predefined rules tutorial. It exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

ifeq ($(shell test -d additional-examples && echo true),true)
		ADDITIONAL_EXAMPLES = $(shell ls -1d ./additional-examples/*/)
else
		ADDITIONAL_EXAMPLES = noop
endif

ADDITIONAL_EXAMPLES_CI := $(patsubst %/,%-ci,$(ADDITIONAL_EXAMPLES))
ADDITIONAL_EXAMPLES_CLEAN := $(patsubst %/,%-clean,$(ADDITIONAL_EXAMPLES))

.PHONY: ci
ci: notice format lint test $(ADDITIONAL_EXAMPLES_CI)

.PHONY: notice
notice:
	@echo "Validating $(realpath ./)"

.PHONY: format
format:
	buf format -w

.PHONY: lint
lint:
	buf lint
	golangci-lint run ./internal/...

.PHONY: test
test:
	go test ./internal/...;

.PHONY: noop
noop:
	@# do nothing

.PHONY: clean
clean:
	git clean -dfX ./

%-ci:
	git clean -dfX $*/ && rsync -r --ignore-existing ./{buf.*,go.*,internal,proto,Makefile} $*/
	cd $*/ && buf generate
	cd $*/ && $(MAKE)

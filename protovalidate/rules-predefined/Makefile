# This Makefile is not intended to be used as part of the Protovalidate
# predefined rules tutorial. It exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

ifeq ($(shell test -d supplementary && echo true),true)
		SUPPLEMENTALS = $(shell ls -1d ./supplementary/*/)
else
		SUPPLEMENTALS = noop
endif

SUPPLEMENTALS_CI := $(patsubst %/,%-ci,$(SUPPLEMENTALS))
SUPPLEMENTALS_CLEAN := $(patsubst %/,%-clean,$(SUPPLEMENTALS))

.PHONY: ci
ci: notice format lint test $(SUPPLEMENTALS_CI)

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

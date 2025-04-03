# This Makefile is not intended to be used as part of Buf
# examples. It exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

BIN := .tmp/bin
COPYRIGHT_HOLDER := Buf Technologies, Inc.
COPYRIGHT_YEARS := 2025
LICENSE_IGNORE := \
	gen
BUF_VERSION := 1.34.0

# Used to transform LICENSE_IGNORE into a comma-delimited list
SPACE :=
SPACE +=
COMMA := ,

.PHONY: all
all: license-header ci

################################################################################
# Code generation/manipulation                                                 #
################################################################################

# Add license headers across all build systems
.PHONY: license-header
license-header: $(BIN)/license-header
	license-header \
		--license-type apache \
		--copyright-holder "$(COPYRIGHT_HOLDER)" \
		--year-range "$(COPYRIGHT_YEARS)" --ignore $(subst $(SPACE),$(COMMA),${LICENSE_IGNORE})


################################################################################
# CI for all examples                                                          #
################################################################################

.PHONY: ci
ci: ci-protovalidate \

.PHONY: ci-protovalidate
ci-protovalidate:
	cd protovalidate && $(MAKE)


################################################################################
# Dependency installation                                                      #
################################################################################

$(BIN)/license-header: Makefile
	@mkdir -p $(@D)
	go install github.com/bufbuild/buf/private/pkg/licenseheader/cmd/license-header@v${BUF_VERSION}

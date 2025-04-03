# This Makefile is not intended to be used as part of the Protovalidate
# quickstart. It exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

BIN := .tmp/bin
PROTOVALIDATE_VERSION := 0.10.3

.PHONY: all
all: clean install ci

.PHONY: clean
clean:
	rm -rf ./vendor/*

################################################################################
# CI for all examples                                                          #
################################################################################

.PHONY: ci
ci: clean \
	install \
 	ci-compiling-buf \
 	ci-compiling-protoc \
 	ci-rules-predefined-finish \
 	ci-rules-predefined-start \
 	ci-quickstart-go-start \
 	ci-quickstart-go-finish \
 	ci-quickstart-java-start \
 	ci-quickstart-java-finish \
 	ci-quickstart-python-start \
 	ci-quickstart-python-finish \
 	ci-connect-go-start \
 	ci-connect-go-finish \
 	ci-grpc-go-start \
 	ci-grpc-go-finish \
 	ci-grpc-java-start \
 	ci-grpc-java-finish \
 	ci-grpc-python-start \
 	ci-grpc-python-finish \
 	ci-bufstream-start \
 	ci-bufstream-finish

.PHONY: ci-rules-predefined-start
ci-rules-predefined-start:
	cd rules-predefined/start&& $(MAKE) ci

.PHONY: ci-rules-predefined-finish
ci-rules-predefined-finish:
	cd rules-predefined/finish&& $(MAKE) ci

.PHONY: ci-compiling-buf
ci-compiling-buf:
	cd compiling-buf&& $(MAKE) ci

.PHONY: ci-compiling-protoc
ci-compiling-protoc:
	cd compiling-protoc&& $(MAKE) ci

.PHONY: ci-quickstart-go-finish
ci-quickstart-go-finish:
	cd quickstart-go/finish&& $(MAKE) ci

.PHONY: ci-quickstart-go-start
ci-quickstart-go-start:
	cd quickstart-go/start&& $(MAKE) ci

.PHONY: ci-quickstart-java-start
ci-quickstart-java-start:
	cd quickstart-java/start&& $(MAKE) ci

.PHONY: ci-quickstart-java-finish
ci-quickstart-java-finish:
	cd quickstart-java/finish&& $(MAKE) ci

.PHONY: ci-quickstart-python-start
ci-quickstart-python-start:
	cd quickstart-python/start && $(MAKE) ci

.PHONY: ci-quickstart-python-finish
ci-quickstart-python-finish:
	cd quickstart-python/finish && $(MAKE) ci

.PHONY: ci-connect-go-start
ci-connect-go-start:
	cd connect-go/start && $(MAKE) ci

.PHONY: ci-connect-go-finish
ci-connect-go-finish:
	cd connect-go/finish && $(MAKE) ci

.PHONY: ci-grpc-java-start
ci-grpc-java-start:
	cd grpc-java/start && $(MAKE) ci

.PHONY: ci-grpc-java-finish
ci-grpc-java-finish:
	cd grpc-java/finish && $(MAKE) ci

.PHONY: ci-grpc-go-start
ci-grpc-go-start:
	cd grpc-go/start && $(MAKE) ci

.PHONY: ci-grpc-go-finish
ci-grpc-go-finish:
	cd grpc-go/finish && $(MAKE) ci

.PHONY: ci-grpc-python-start
ci-grpc-python-start:
	cd grpc-python/start && $(MAKE) ci

.PHONY: ci-grpc-python-finish
ci-grpc-python-finish:
	cd grpc-python/finish && $(MAKE) ci

.PHONY: ci-bufstream-start
ci-bufstream-start:
	cd bufstream/start && $(MAKE) ci

.PHONY: ci-bufstream-finish
ci-bufstream-finish:
	cd bufstream/finish && $(MAKE) ci

################################################################################
# Dependency installation                                                      #
################################################################################

install:
	# Use the declared version of Protovalidate to download one validate.proto
	# to use across any examples that require it. Examples are responsible for
	# installing from this source of truth and should not download their own.
	curl --create-dirs -f \
		-O \
		--output-dir ./vendor/github.com/bufbuild/protovalidate/buf/validate \
		https://raw.githubusercontent.com/bufbuild/protovalidate/refs/tags/v${PROTOVALIDATE_VERSION}/proto/protovalidate/buf/validate/validate.proto

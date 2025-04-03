# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := build
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

EXPECTED_FILE := gen/proto/example/v1/building_example.pb.cc

# Targets for compilation / generation

.PHONY: build
build:
	mkdir -p ./gen
	protoc ./proto/example/v1/building_example.proto \
		-I. -I../vendor/github.com/bufbuild/protovalidate/ \
		--cpp_out=:./gen

# Targets for buf-examples repository CI purposes

.PHONY: clean
clean:
	rm -rf gen

.PHONY: format
format:
	buf format -w

.PHONY: ci
ci: format clean build
	ls $(EXPECTED_FILE)

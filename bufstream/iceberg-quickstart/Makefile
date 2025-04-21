# This Makefile is not intended to be used as part of the Iceberg
# quickstart. It exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.PHONY: clean
clean: down
	rm -rf minio/*
	rm -rf postgres/*
	rm -rf iceberg-rest/*

.PHONY: up
up:
	docker compose up -d

.PHONY: stop
stop:
	docker compose stop

.PHONY: down
down:
	docker compose down

.PHONY: publish
publish:
	go run ./cmd/producer \
      --topic email-updated \
      --group email-verifier \
      --csr-url "https://demo.buf.dev/integrations/confluent/bufstream-demo"

.PHONY: archive
archive:
	docker exec bufstream /usr/local/bin/bufstream admin clean topics

.PHONY: ci
ci: format-proto lint-proto

.PHONY: format-proto
format-proto:
	buf format -w

.PHONY: lint-proto
lint-proto:
	buf lint

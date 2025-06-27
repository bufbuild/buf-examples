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
clean:
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
down: docker-down clean

.PHONY: docker-down
docker-down:
	docker compose down

.PHONY: configure
configure:
	go run ./cmd/configure \
		--topic email-updated \
		--bufstream.archive.kind ICEBERG \
		--bufstream.archive.iceberg.catalog local-rest-catalog \
		--bufstream.archive.iceberg.table bufstream.email_updated

.PHONY: produce
produce:
	go run ./cmd/produce \
      --topic email-updated \
      --csr-url "https://demo.buf.dev/integrations/confluent/bufstream-demo"

.PHONY: archive
archive:
	docker exec bufstream /usr/local/bin/bufstream admin clean topics

.PHONY: add-records
add-records: produce archive

.PHONY: reset-table
reset-table: down up configure add-records
	@echo "Open your browser to http://localhost:8888/notebooks/notebooks/bufstream-iceberg-quickstart.ipynb to continue"

.PHONY: ci
ci: format-proto lint-proto lint-go

.PHONY: format-proto
format-proto:
	buf format -w

.PHONY: lint-proto
lint-proto:
	buf lint

.PHONY: lint-go
lint-go:
	golangci-lint run ./...

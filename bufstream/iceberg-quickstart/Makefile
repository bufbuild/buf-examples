# This Makefile is not intended to be used as part of the Iceberg
# quickstart. It exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

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

# Convenience targets for running this as a demo.
.PHONY: run
run: docker-up create-topic produce archive
	echo "Open http://localhost:8888/notebooks/notebooks/bufstream-iceberg-quickstart.ipynb to query the bufstream.email_updated table."

.PHONY: docker-down
docker-down:
	docker compose down

.PHONY: docker-up
docker-up: docker-down
	docker compose up -d

.PHONY: create-topic
create-topic:
	./create-topic.sh

.PHONY: produce
produce:
	./produce.sh

.PHONY: archive
archive:
	docker exec bufstream /usr/local/bin/bufstream admin clean topics


# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := broker
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

.DEFAULT_GOAL := broker

.PHONY: clean
clean:
	rm -f bufstream

.PHONY: broker
broker: bufstream
	./bufstream \
		serve \
		--config config/bufstream.yaml

bufstream:
	curl -sSL -o bufstream \
        "https://buf.build/dl/bufstream/latest/bufstream-$(shell uname -s)-$(shell uname -m)" && \
        chmod +x bufstream

.PHONY: consumer
consumer:
	go run ./cmd/bufstream-demo-consume \
	  --topic email-updated \
	  --group email-verifier \
	  --csr-url "https://demo.buf.dev/integrations/confluent/bufstream-demo"

.PHONY: producer
producer:
	go run ./cmd/bufstream-demo-produce \
	  --topic email-updated \
	  --group email-verifier \
	  --csr-url "https://demo.buf.dev/integrations/confluent/bufstream-demo"

.PHONY: ci
ci: format-proto lint-proto

.PHONY: format-proto
format-proto:
	buf format -w

.PHONY: lint-proto
lint-proto:
	buf lint

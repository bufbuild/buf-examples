# This Makefile is not intended to be used as part of the Protovalidate
# quickstart. It exists only for CI purposes.

# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := ci
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

VENV = .venv
PYTHON = $(VENV)/bin/python3
PIP = $(VENV)/bin/pip

venv: clean
	@echo "Creating a new virtual environment..."
	python3 -m venv $(VENV)
	$(PIP) install -r requirements.txt --extra-index-url https://buf.build/gen/python
	touch $(VENV)/bin/activate
	@echo "Virtual environment rejuvenated."

clean:
	rm -rf __pycache__
	rm -rf $(VENV)

.PHONY: ci
ci: venv format-proto lint-proto lint-python generate test

.PHONY: format-proto
format-proto:
	buf format -w

.PHONY: lint-proto
lint-proto:
	buf lint

.PHONY: lint-python
lint-python:
	ruff check weather_test.py

.PHONY: generate
generate:
	buf generate

.PHONY: test
test:
	$(PYTHON) -m unittest -v weather_test.py

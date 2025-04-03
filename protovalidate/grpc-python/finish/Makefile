# Note: this Makefile exists for CI purposes. See Protovalidate's documentation
# at http://buf.build/docs/protovalidate for how to use this tutorial.

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

.PHONY: clean
clean:
	rm -rf __pycache__
	rm -rf $(VENV)

.PHONY: ci
ci: clean venv format-proto lint-proto lint-python test

.PHONY: format-proto
format-proto:
	buf format -w

.PHONY: lint-proto
lint-proto:
	buf lint

.PHONY: lint-python
lint-python:
	ruff check invoice_server.py validation/*

.PHONY: test
test:
	$(PYTHON) -m unittest -v invoice_server_test.py

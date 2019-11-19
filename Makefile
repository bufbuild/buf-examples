SHELL := /usr/bin/env bash -o pipefail

# This controls the location of the cache.
PROJECT := buf-example
# This controls the remote git location to compare against for breaking changes in CI
#
# Most CI providers only clone the branch under test and to a certain depth, so when
# running buf check breaking in CI, it is generally preferable to compare against
# the remote repository directly.
REMOTE_GIT := https://github.com/bufbuild/buf-example.git
# This controls the version of buf to install and use.
BUF_VERSION := 0.3.0

### Everything below this line is meant to be static, i.e. only adjust the above variables. ###

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
# Buf will be cached to ~/.cache/buf-example.
CACHE_BASE := $(HOME)/.cache/$(PROJECT)
# This allows switching between i.e a Docker container and your local setup without overwriting.
CACHE := $(CACHE_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
# The location where buf will be installed.
CACHE_BIN := $(CACHE)/bin
# Marker files are put into this directory to denote the current version of binaries that are installed.
CACHE_VERSIONS := $(CACHE)/versions

# Update the $PATH so we can use buf directly
export PATH := $(abspath $(CACHE_BIN)):$(PATH)

# BUF points to the marker file for the installed version.
#
# If BUF_VERSION is changed, the binary will be re-downloaded.
BUF := $(CACHE_VERSIONS)/buf/$(BUF_VERSION)
$(BUF):
	@rm -f $(CACHE_BIN)/buf
	@mkdir -p $(CACHE_BIN)
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/buf-$(UNAME_OS)-$(UNAME_ARCH)" \
		-o "$(CACHE_BIN)/buf"
	chmod +x "$(CACHE_BIN)/buf"
	@rm -rf $(dir $(BUF))
	@mkdir -p $(dir $(BUF))
	@touch $(BUF)

.DEFAULT_GOAL := local

# deps allows us to install deps without running any checks.

.PHONY: deps
deps: $(BUF)

# local is what we run when testing locally
# this does breaking change detection against our local git repository

.PHONY: local
local: $(BUF)
	buf check lint
	buf check breaking --against-input '.git#branch=master'

# remote is what we run when testing in most CI providers
# this does breaking change detection against our remote git repository
#
.PHONY: remote
remote: $(BUF)
	buf check lint
	buf check breaking --against-input "$(REMOTE_GIT)#branch=master"

# clean deletes any files not checked in and the cache for all platforms

.PHONY: clean
clean:
	git clean -xdf
	rm -rf $(CACHE_BASE)

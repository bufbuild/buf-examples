BUF_VERSION := 1.0.0
UNAME_OS := $(shell uname -s)
FILES_TO_UPDATE := Makefile README.md .circleci/config.yml .github/workflows/ci.yaml

ifeq ($(UNAME_OS),Darwin)
SED_I := sed -i ''
else
SED_I := sed -i
endif

.PHONY: updateversion
updateversion:
ifndef VERSION
	$(error "VERSION must be set")
else
	$(foreach file, $(FILES_TO_UPDATE), $(SED_I) "s/BUF_VERSION := [0-9][0-9]*\.[0-9][0-9]*\.[0-9][0-9]*/BUF_VERSION := $(VERSION)/g" $(file))
endif

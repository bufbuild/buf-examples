PREFIX         =  plugins.buf.build
GO_VERSION     =  1.17
OWNER          =  change-this
NAME           =  change-this-too
PLUGIN         = $(OWNER)/$(NAME)
PLUGIN_VERSION =  v0.1.0
TAG            := $(PREFIX)/$(PLUGIN):$(PLUGIN_VERSION)

build:
	docker build \
	--tag $(TAG) \
	--build-arg goVersion=$(GO_VERSION) \
	--build-arg pluginOwner=$(OWNER) \
	--build-arg pluginName=$(NAME) \
	--build-arg pluginVersion=$(PLUGIN_VERSION) \
	$(CURDIR)

push:
	docker push $(TAG)

generate:
	go install ./cmd/protoc-gen-go-example
	protoc \
	--go-example_out=proto \
	proto/testing.proto

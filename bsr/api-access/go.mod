module github.com/bufbuild/buf-examples/bsr/api-access

go 1.26.0

require (
	buf.build/gen/go/bufbuild/reflect/connectrpc/go/v2 v2.0.2-20240117202343-bf8f65e8876c.1
	buf.build/gen/go/bufbuild/reflect/protocolbuffers/go v1.36.12-20240117202343-bf8f65e8876c.2
	connectrpc.com/connect/v2 v2.0.2
	github.com/stretchr/testify v1.12.1
)

require (
	go.yaml.in/yaml/v3 v3.0.6 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

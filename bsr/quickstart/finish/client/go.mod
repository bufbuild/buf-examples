module github.com/bufbuild/buf-examples/bsr/quickstart/client

go 1.24.0

require (
	buf.build/gen/go/xUSERNAMEx/common/protocolbuffers/go v1.36.5-20250319033322-7607f73aaca1.1
	buf.build/gen/go/quickstarts/bsr/connectrpc/go v1.18.1-20250319031420-b7f0cadb0629.1
	buf.build/gen/go/quickstarts/bsr/protocolbuffers/go v1.36.5-20250319031420-b7f0cadb0629.1
	connectrpc.com/connect v1.18.1
)

require google.golang.org/protobuf v1.36.5 // indirect

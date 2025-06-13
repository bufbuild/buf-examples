module github.com/bufbuild/buf-examples

go 1.23.0

require (
	connectrpc.com/connect v1.16.1
	golang.org/x/net v0.38.0
	google.golang.org/protobuf v1.34.2
)

require golang.org/x/text v0.23.0 // indirect

// Temporary! The connect change is pending. This path works on Joe's laptop.
replace connectrpc.com/connect => ../../../../../connectrpc/connect-go

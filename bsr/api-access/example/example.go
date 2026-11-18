package example

import (
	"context"
	"net/http"

	reflectv1beta1connect "buf.build/gen/go/bufbuild/reflect/connectrpc/go/v2/buf/reflect/v1beta1/reflectv1beta1connect"
	reflectv1beta1 "buf.build/gen/go/bufbuild/reflect/protocolbuffers/go/buf/reflect/v1beta1"
	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
)

func Example() (*reflectv1beta1.GetFileDescriptorSetResponse, error) {
	client := reflectv1beta1connect.NewFileDescriptorSetServiceClient(connect.NewClient(
		connecthttp.NewTransport(http.DefaultClient, "https://buf.build"),
	))
	request := &reflectv1beta1.GetFileDescriptorSetRequest{
		Module: "buf.build/connectrpc/eliza",
	}
	ctx := context.Background()
	// If you're using a private BSR, set your Authorization header to a
	// BUF_TOKEN value.
	//
	// ctx, info := connect.NewClientContext(ctx)
	// info.RequestHeader().Set("Authorization", "Bearer <BUF_TOKEN>")
	response, err := client.GetFileDescriptorSet(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

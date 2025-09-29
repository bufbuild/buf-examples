package example

import (
	"context"
	"net/http"

	reflectv1beta1connect "buf.build/gen/go/bufbuild/reflect/connectrpc/gosimple/buf/reflect/v1beta1/reflectv1beta1connect"
	reflectv1beta1 "buf.build/gen/go/bufbuild/reflect/protocolbuffers/go/buf/reflect/v1beta1"
)

func Example() (*reflectv1beta1.GetFileDescriptorSetResponse, error) {
	client := reflectv1beta1connect.NewFileDescriptorSetServiceClient(
		http.DefaultClient,
		"https://buf.build",
	)
	request := &reflectv1beta1.GetFileDescriptorSetRequest{
		Module: "buf.build/connectrpc/eliza",
	}
	// If you're using a private BSR, set your Authorization header to a
	// BUF_TOKEN value.
	//
	// request.Header().Set("Authorization", "Bearer <BUF_TOKEN>")
	response, err := client.GetFileDescriptorSet(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

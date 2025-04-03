package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	connect "connectrpc.com/connect"
	petv1 "github.com/bufbuild/buf-tour/gen/pet/v1"
	"github.com/bufbuild/buf-tour/gen/pet/v1/petv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(&petStoreServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("... Listening on", address)
	http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

// petStoreServiceServer implements the PetStoreService API.
type petStoreServiceServer struct {
	petv1connect.UnimplementedPetStoreServiceHandler
}

// PutPet adds the pet associated with the given request into the PetStore.
func (s *petStoreServiceServer) PutPet(
	ctx context.Context,
	req *connect.Request[petv1.PutPetRequest],
) (*connect.Response[petv1.PutPetResponse], error) {
	name := req.Msg.GetName()
	petType := req.Msg.GetPetType()
	log.Printf("Got a request to create a %v named %s", petType, name)
	return connect.NewResponse(&petv1.PutPetResponse{}), nil
}

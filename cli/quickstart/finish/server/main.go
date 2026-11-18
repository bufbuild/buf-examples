package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	petv1 "github.com/bufbuild/buf-examples/gen/pet/v1"
	"github.com/bufbuild/buf-examples/gen/pet/v1/petv1connect"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(&petStoreServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("... Listening on", address)
	// Serve HTTP/2 without TLS (h2c) alongside HTTP/1.1.
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)
	server := &http.Server{
		Addr:      address,
		Handler:   mux,
		Protocols: protocols,
	}
	server.ListenAndServe()
}

// petStoreServiceServer implements the PetStoreService API.
type petStoreServiceServer struct {
	petv1connect.UnimplementedPetStoreServiceHandler
}

// PutPet adds the pet associated with the given request into the PetStore.
func (s *petStoreServiceServer) PutPet(
	_ context.Context,
	req *petv1.PutPetRequest,
) (*petv1.PutPetResponse, error) {
	pet := &petv1.Pet{Name: req.GetName(), PetType: req.PetType}
	log.Printf("PutPet received a %v named %s", pet.GetPetType(), pet.GetName())
	return &petv1.PutPetResponse{Pet: pet}, nil
}

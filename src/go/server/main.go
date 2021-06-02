package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	duration "github.com/golang/protobuf/ptypes/duration"
	petv1 "go.bufbuild.internal/local/go-grpc/bufexample/petapis/bufexample/pet/v1"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PetStoreServer struct {
	petv1.UnimplementedPetStoreServer

	petsMu sync.Mutex
	pets   map[string]*petv1.Pet
}

func (s *PetStoreServer) GetPet(ctx context.Context, req *petv1.GetPetRequest) (*petv1.GetPetResponse, error) {
	id := req.GetPetId()
	if len(id) < 1 {
		return nil, status.Error(codes.InvalidArgument, "pet id must not be empty")
	}

	s.petsMu.Lock()
	defer s.petsMu.Unlock()
	pet, ok := s.pets[id]
	if !ok {
		return nil, status.Error(codes.NotFound, "pet not found")
	}

	fmt.Printf("Retrieved ID %s whose name is %s\n", id, pet.GetName())
	return &petv1.GetPetResponse{
		Pet: pet,
	}, nil
}

func (s *PetStoreServer) PutPet(ctx context.Context, req *petv1.PutPetRequest) (*petv1.PutPetResponse, error) {
	name := req.GetName()
	if len(name) < 1 {
		return nil, status.Error(codes.InvalidArgument, "pet name must not be empty")
	}
	petType := req.GetPetType()
	if petType == petv1.PetType_PET_TYPE_UNSPECIFIED {
		return nil, status.Error(codes.InvalidArgument, "pet type must be specified")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate pet ID: %w", err)
	}

	s.petsMu.Lock()
	defer s.petsMu.Unlock()
	p := &petv1.Pet{
		PetType:   petType,
		PetId:     id.String(),
		Name:      name,
		CreatedAt: datetimeNow(),
	}
	s.pets[id.String()] = p

	fmt.Printf("Created %s who is a %s with ID %s\n", p.GetName(), p.GetPetType().String(), id.String())
	return &petv1.PutPetResponse{
		Pet: p,
	}, nil
}

func (s *PetStoreServer) DeletePet(ctx context.Context, req *petv1.DeletePetRequest) (*petv1.DeletePetResponse, error) {
	id := req.GetPetId()
	if len(id) < 1 {
		return nil, status.Error(codes.InvalidArgument, "pet id must not be empty")
	}

	s.petsMu.Lock()
	defer s.petsMu.Unlock()
	if _, ok := s.pets[id]; !ok {
		return nil, status.Error(codes.NotFound, "pet id not found")
	}
	delete(s.pets, id)

	return &petv1.DeletePetResponse{}, nil
}

func (s *PetStoreServer) PurchasePet(ctx context.Context, req *petv1.PurchasePetRequest) (*petv1.PurchasePetResponse, error) {
	id := req.GetPetId()
	if len(id) < 1 {
		return nil, status.Error(codes.InvalidArgument, "pet id must not be empty")
	}
	order := req.GetOrder()
	if order == nil {
		return nil, status.Error(codes.InvalidArgument, "order must be provided")
	}
	amount := order.GetAmount()
	if amount == nil {
		return nil, status.Error(codes.InvalidArgument, "an amount must be provided")
	}

	s.petsMu.Lock()
	defer s.petsMu.Unlock()
	pet, ok := s.pets[id]
	if !ok {
		return nil, status.Error(codes.NotFound, "pet id not found")
	}

	wholeUnits := amount.GetUnits()
	totalUnits := float64(wholeUnits) + float64(amount.GetNanos())/1_000_000_000
	fmt.Printf("Purchased %s who is a %s for %f %s\n", pet.GetName(), pet.GetPetType().String(), totalUnits, amount.GetCurrencyCode())
	return &petv1.PurchasePetResponse{}, nil
}

func datetimeNow() *datetime.DateTime {
	now := time.Now()
	return &datetime.DateTime{
		Year:    int32(now.Year()),
		Month:   int32(now.Month()),
		Day:     int32(now.Day()),
		Hours:   int32(now.Hour()),
		Minutes: int32(now.Minute()),
		Seconds: int32(now.Second()),
		Nanos:   int32(now.Nanosecond()),
		TimeOffset: &datetime.DateTime_UtcOffset{
			UtcOffset: &duration.Duration{
				Seconds: 0,
				Nanos:   0,
			},
		},
	}
}

func main() {
	listenOn := "127.0.0.1:1337"
	ln, err := net.Listen("tcp", listenOn)
	if err != nil {
		fmt.Printf("failed to listen on %s: %v\n", listenOn, err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	petv1.RegisterPetStoreServer(grpcServer, &PetStoreServer{
		pets: make(map[string]*petv1.Pet),
	})

	fmt.Printf("Listening on %s\n", listenOn)
	if err := grpcServer.Serve(ln); err != nil {
		fmt.Printf("failed to serve grpc server: %v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
	paymentv1 "go.bufbuild.internal/local/go-grpc/bufexample/paymentapis/bufexample/payment/v1"
	petv1 "go.bufbuild.internal/local/go-grpc/bufexample/petapis/bufexample/pet/v1"
	"google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/grpc"
)

func newPutCmd(client petv1.PetStoreClient) *ffcli.Command {
	fs := flag.NewFlagSet("put", flag.ExitOnError)
	name := fs.String("name", "", "Pet Name")
	petType := fs.String("type", "", "Pet Type")
	return &ffcli.Command{
		Name:       "put",
		ShortUsage: "client put -name <name> -type <type>",
		ShortHelp:  "Put a new pet into the store",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			if nargs := len(args); nargs > 0 {
				return fmt.Errorf("put only takes flag arguments, you provided %d positional arguments", nargs)
			}
			if len(*name) < 1 {
				return errors.New("put requires a pet name to be specified")
			}
			petTypeEnumInt, ok := petv1.PetType_value["PET_TYPE_"+strings.ToUpper(*petType)]
			if !ok {
				return fmt.Errorf("%q is not a valid pet type", *petType)
			}
			petTypeEnum := petv1.PetType(petTypeEnumInt)

			petResp, err := client.PutPet(ctx, &petv1.PutPetRequest{
				Name:    *name,
				PetType: petTypeEnum,
			})
			if err != nil {
				return fmt.Errorf("failed to create pet: %w", err)
			}
			pet := petResp.GetPet()
			fmt.Printf("Created a %s named %s with ID %s\n", pet.GetPetType().String(), pet.GetName(), pet.GetPetId())
			return nil
		},
	}
}

func newGetCmd(client petv1.PetStoreClient) *ffcli.Command {
	fs := flag.NewFlagSet("get", flag.ExitOnError)
	id := fs.String("id", "", "Pet ID")
	return &ffcli.Command{
		Name:       "get",
		ShortUsage: "client get -id <id>",
		ShortHelp:  "Get a pet from the store",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			if nargs := len(args); nargs > 0 {
				return fmt.Errorf("get only takes flag arguments, you provided %d positional arguments", nargs)
			}
			if len(*id) < 1 {
				return errors.New("get requires a pet id to be specified")
			}

			petResp, err := client.GetPet(ctx, &petv1.GetPetRequest{
				PetId: *id,
			})
			if err != nil {
				return fmt.Errorf("failed to get pet: %w", err)
			}
			pet := petResp.GetPet()
			fmt.Printf("Got a %s named %s with ID %s\n", pet.GetPetType().String(), pet.GetName(), pet.GetPetId())
			return nil
		},
	}
}

func newDeleteCmd(client petv1.PetStoreClient) *ffcli.Command {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	id := fs.String("id", "", "Pet ID")
	return &ffcli.Command{
		Name:       "delete",
		ShortUsage: "client delete -id <id>",
		ShortHelp:  "delete a pet from the store",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			if nargs := len(args); nargs > 0 {
				return fmt.Errorf("delete only takes flag arguments, you provided %d positional arguments", nargs)
			}
			if len(*id) < 1 {
				return errors.New("delete requires a pet id to be specified")
			}

			_, err := client.DeletePet(ctx, &petv1.DeletePetRequest{
				PetId: *id,
			})
			if err != nil {
				return fmt.Errorf("failed to delete pet: %w", err)
			}
			fmt.Printf("Deleted pet with ID %s\n", *id)
			return nil
		},
	}
}

func newPurchaseCmd(client petv1.PetStoreClient) *ffcli.Command {
	fs := flag.NewFlagSet("purchase", flag.ExitOnError)
	id := fs.String("id", "", "Pet ID")
	currencyCode := fs.String("currency", "", "Purchase currency code, 3 letters")
	currencyAmount := fs.String("value", "", "Purchase value including fractional units e.g. 123.45")
	return &ffcli.Command{
		Name:       "purchase",
		ShortUsage: "client purchase -id <id> -currency <currency> -value <value>",
		ShortHelp:  "purchase a pet from the store",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			if nargs := len(args); nargs > 0 {
				return fmt.Errorf("purchase only takes flag arguments, you provided %d positional arguments", nargs)
			}
			if len(*id) < 1 {
				return errors.New("purchase requires a pet id to be specified")
			}
			if curLen := len(*currencyCode); curLen != 3 {
				return fmt.Errorf("currency code must be exactly 3 letters, got %q", *currencyCode)
			}
			amountAsFloat, err := strconv.ParseFloat(*currencyAmount, 64)
			if err != nil {
				return fmt.Errorf("unable to convert value %q into a float: %w", *currencyAmount, err)
			}
			units, remainderAsFloat := math.Modf(amountAsFloat)
			nanoUnits := remainderAsFloat * 1_000_000_000

			_, err = client.PurchasePet(ctx, &petv1.PurchasePetRequest{
				PetId: *id,
				Order: &paymentv1.Order{
					Amount: &money.Money{
						CurrencyCode: *currencyCode,
						Units:        int64(units),
						Nanos:        int32(nanoUnits),
					},
				},
			})
			if err != nil {
				return fmt.Errorf("failed to delete pet: %w", err)
			}
			fmt.Printf("Purchased pet with ID %s for %f %s\n", *id, amountAsFloat, *currencyCode)
			return nil
		},
	}
}

func main() {
	connectTo := "127.0.0.1:1337"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Could not connect to PetStore server on %s: %v\n", connectTo, err)
		os.Exit(1)
	}
	client := petv1.NewPetStoreClient(conn)

	putCmd := newPutCmd(client)
	getCmd := newGetCmd(client)
	deleteCmd := newDeleteCmd(client)
	purchaseCmd := newPurchaseCmd(client)
	rootCmd := &ffcli.Command{
		Subcommands: []*ffcli.Command{
			getCmd,
			putCmd,
			deleteCmd,
			purchaseCmd,
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}

	if err := rootCmd.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

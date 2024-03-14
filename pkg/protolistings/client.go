package protolistings

import (
	"log"

	"github.com/pandishpancheta/contract-observer-service/pkg/config"
	listings "github.com/pandishpancheta/contract-observer-service/pkg/protolistings/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client listings.ListingsServiceClient
}

func InitServiceClient(cfg *config.Config) (listings.ListingsServiceClient, error) {
	creds := insecure.NewCredentials()

	c, err := grpc.Dial(cfg.ListingsServiceClient, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to dial TokenizationService: %v", err)
		return nil, err
	}

	return listings.NewListingsServiceClient(c), nil
}

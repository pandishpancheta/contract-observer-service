package listings

import (
	"log"

	"github.com/pandishpanceta/contract-observer-service/pkg/config"
	"github.com/pandishpanceta/contract-observer-service/pkg/listings/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ListingsServiceClient
}

func InitServiceClient(cfg *config.Config) (pb.ListingsServiceClient, error) {
	creds := insecure.NewCredentials()

	c, err := grpc.Dial(cfg.ListingsServiceClient, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to dial TokenizationService: %v", err)
		return nil, err
	}

	return pb.NewListingsServiceClient(c), nil
}

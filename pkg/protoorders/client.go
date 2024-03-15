package protoorders

import (
	"log"

	"github.com/pandishpancheta/contract-observer-service/pkg/config"
	orders "github.com/pandishpancheta/contract-observer-service/pkg/protoorders/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client orders.OrderServiceClient
}

func InitServiceClient(cfg *config.Config) (orders.OrderServiceClient, error) {
	creds := insecure.NewCredentials()

	c, err := grpc.Dial(cfg.OrdersServiceClient, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to dial TokenizationService: %v", err)
		return nil, err
	}

	return orders.NewOrderServiceClient(c), nil
}

package config

import (
	"os"
	"strings"
)

type Config struct {
	WSUrl                 string
	ContractAddress       string
	ListingsServiceClient string
	OrdersServiceClient   string
}

func LoadConfig() Config {
	return Config{
		WSUrl:                 strings.TrimSpace(os.Getenv("WS_URL")),
		ContractAddress:       strings.TrimSpace(os.Getenv("CONTRACT_ADDRESS")),
		ListingsServiceClient: strings.TrimSpace(os.Getenv("LISTINGS_SERVICE_CLIENT")),
		OrdersServiceClient:   strings.TrimSpace(os.Getenv("ORDERS_SERVICE_CLIENT")),
	}
}

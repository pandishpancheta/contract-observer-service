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
		WSUrl:                 strings.TrimSuffix(os.Getenv("WS_URL"), "\n"),
		ContractAddress:       strings.TrimSuffix(os.Getenv("CONTRACT_ADDRESS"), "\n"),
		ListingsServiceClient: strings.TrimSuffix(os.Getenv("LISTINGS_SERVICE_CLIENT"), "\n"),
		OrdersServiceClient:   strings.TrimSuffix(os.Getenv("ORDERS_SERVICE_CLIENT"), "\n"),
	}
}

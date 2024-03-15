package config

import "os"

type Config struct {
	WSUrl                 string
	ContractAddress       string
	ListingsServiceClient string
	OrdersServiceClient   string
}

func LoadConfig() Config {
	return Config{
		WSUrl:                 os.Getenv("WS_URL"),
		ContractAddress:       os.Getenv("CONTRACT_ADDRESS"),
		ListingsServiceClient: os.Getenv("LISTINGS_SERVICE_CLIENT"),
		OrdersServiceClient:   os.Getenv("ORDERS_SERVICE_CLIENT"),
	}
}

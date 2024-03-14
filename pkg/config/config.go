package config

import "os"

type Config struct {
	WSUrl                 string
	ContractAddress       string
	ListingsServiceClient string
}

func LoadConfig() Config {
	return Config{
		WSUrl:                 os.Getenv("WS_URL"),
		ContractAddress:       os.Getenv("CONTRACT_ADDRESS"),
		ListingsServiceClient: os.Getenv("LISTINGS_SERVICE_CLIENT"),
	}
}

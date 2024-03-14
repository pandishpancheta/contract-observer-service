package config

import "os"

type Config struct {
	WSUrl           string
	ContractAddress string
}

func LoadConfig() Config {
	return Config{
		WSUrl:           os.Getenv("WS_URL"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
	}
}

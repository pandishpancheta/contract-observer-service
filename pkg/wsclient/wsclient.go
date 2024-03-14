package wsclient

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pandishpancheta/contract-observer-service/pkg/config"
)

func InitWsClient(cfg *config.Config) (*ethclient.Client, error) {
	client, err := ethclient.Dial(cfg.WSUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}

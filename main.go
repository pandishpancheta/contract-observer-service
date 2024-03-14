package main

import (
	"log"

	"github.com/pandishpanceta/contract-observer-service/pkg/config"
	"github.com/pandishpanceta/contract-observer-service/pkg/events"
)

func main() {
	log.Println("Loading config...")
	cfg := config.LoadConfig()

	log.Println("Starting contract observer service...")
	events.RunSubscription(&cfg, "./pkg/events/abi.json")
}

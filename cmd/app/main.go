package main

import (
	"fmt"
	"katify/internal/adapters/server"
	"katify/internal/config"
	"katify/pkg/logger"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to set up the config: %w", err))
	}
	log := logger.NewLogger(*cfg.Logger)
	log.Fatalf("Bye bye", server.NewServer(cfg, log).Run())
}

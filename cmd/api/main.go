package main

import (
	"log"

	"github.com/alves/age-of-genesis/internal/config"
	"github.com/alves/age-of-genesis/internal/interfaces/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server, err := http.NewServer(cfg)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}

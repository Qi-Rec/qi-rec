package main

import (
	"log"

	"qi-rec/internal/app"
	"qi-rec/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	a := app.New(cfg)

	a.Run()
}
